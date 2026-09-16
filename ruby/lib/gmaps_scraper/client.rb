# frozen_string_literal: true

require "json"
require "net/http"
require "uri"

module GmapsScraper
  class Client
    DEFAULT_BASE_URL = "https://gmapsleadfinder.com"
    DEFAULT_POLL_INTERVAL_MS = 2000
    DEFAULT_TIMEOUT_MS = 600_000
    DEFAULT_RESULT_LIMIT = 100
    DEFAULT_HTTP_TIMEOUT_S = 60.0
    USER_AGENT = "gmaps-scraper-ruby/#{VERSION}"
    TERMINAL_STATUSES = %w[completed partial failed].freeze

    def initialize(api_key: nil, base_url: nil, timeout_s: DEFAULT_HTTP_TIMEOUT_S)
      key = (api_key || ENV["GMF_API_KEY"]).to_s.strip
      if key.empty?
        raise AuthenticationError,
              "Missing API key. Set GMF_API_KEY or pass api_key: " \
              "(https://gmapsleadfinder.com/account#api-key)"
      end

      @api_key = key
      @base_url = (base_url || ENV["GMF_BASE_URL"] || DEFAULT_BASE_URL).to_s.strip.gsub(%r{/+\z}, "")
      @timeout_s = timeout_s
    end

    def me
      request("GET", "/api/v1/me")
    end

    def create_job(keyword)
      trimmed = keyword.to_s.strip
      raise BadRequestError, "keyword is required" if trimmed.empty?

      request("POST", "/api/v1/jobs", json_body: { "keyword" => trimmed })
    end

    def get_job(job_id)
      request("GET", "/api/v1/jobs/#{job_id}")
    end

    def get_results(job_id, limit: DEFAULT_RESULT_LIMIT, cursor: nil)
      query = { "limit" => limit.to_s }
      query["cursor"] = cursor.to_s unless cursor.nil?
      request("GET", "/api/v1/jobs/#{job_id}/results", query: query)
    end

    def scrape(
      keyword,
      poll_interval_ms: DEFAULT_POLL_INTERVAL_MS,
      timeout_ms: DEFAULT_TIMEOUT_MS,
      result_limit: DEFAULT_RESULT_LIMIT
    )
      created = create_job(keyword)
      job_id = created["jobId"].to_s
      deadline = Process.clock_gettime(Process::CLOCK_MONOTONIC) + (timeout_ms / 1000.0)
      poll_s = [poll_interval_ms, 100].max / 1000.0

      loop do
        job = get_job(job_id)
        status = job["status"].to_s.downcase
        if TERMINAL_STATUSES.include?(status)
          if status == "failed"
            raise ApiError.new(job["error"] || "Job #{job_id} failed", body: job)
          end
          break
        end
        if Process.clock_gettime(Process::CLOCK_MONOTONIC) >= deadline
          raise TimeoutError,
                "Timed out after #{timeout_ms}ms waiting for job #{job_id} " \
                "(last status=#{job['status'].inspect})"
        end
        sleep(poll_s)
      end

      fetch_all_rows(job_id, result_limit: result_limit)
    end

    private

    def fetch_all_rows(job_id, result_limit:)
      rows = []
      cursor = "0"
      while cursor
        page = get_results(job_id, limit: result_limit, cursor: cursor)
        rows.concat(Array(page["rows"]))
        next_cursor = page["nextCursor"]
        cursor = next_cursor.nil? || next_cursor == "" ? nil : next_cursor.to_s
      end
      rows
    end

    def request(method, path, json_body: nil, query: nil)
      uri = URI("#{@base_url}#{path}")
      if query && !query.empty?
        uri.query = URI.encode_www_form(query)
      end

      http = Net::HTTP.new(uri.host, uri.port)
      http.use_ssl = uri.scheme == "https"
      http.open_timeout = @timeout_s
      http.read_timeout = @timeout_s

      klass = case method.upcase
              when "GET" then Net::HTTP::Get
              when "POST" then Net::HTTP::Post
              when "PUT" then Net::HTTP::Put
              when "DELETE" then Net::HTTP::Delete
              else
                raise ApiError, "Unsupported HTTP method #{method}"
              end

      req = klass.new(uri)
      req["Authorization"] = "Bearer #{@api_key}"
      req["Accept"] = "application/json"
      req["User-Agent"] = USER_AGENT
      if json_body
        req["Content-Type"] = "application/json"
        req.body = JSON.generate(json_body)
      end

      begin
        res = http.request(req)
      rescue StandardError => e
        raise ApiError, "Network error: #{e.message}"
      end

      body = parse_body(res.body)
      code = res.code.to_i
      raise_for_status(code, body) if code < 200 || code >= 300
      body.is_a?(Hash) ? body : {}
    end

    def parse_body(raw)
      return {} if raw.nil? || raw.empty?

      JSON.parse(raw)
    rescue JSON::ParserError
      { "message" => raw }
    end

    def raise_for_status(status, body)
      message = "API request failed with status #{status}"
      if body.is_a?(Hash)
        %w[message error statusMessage].each do |key|
          val = body[key]
          if val.is_a?(String) && !val.empty?
            message = val
            break
          end
        end
      end

      klass = case status
              when 400 then BadRequestError
              when 401 then AuthenticationError
              when 402 then InsufficientCreditsError
              when 403 then PlanNotAllowedError
              when 404 then NotFoundError
              when 409 then JobConflictError
              else ApiError
              end
      raise klass.new(message, status_code: status, body: body)
    end
  end
end
