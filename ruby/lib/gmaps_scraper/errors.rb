# frozen_string_literal: true

module GmapsScraper
  class ApiError < StandardError
    attr_reader :status_code, :body

    def initialize(message, status_code: nil, body: nil)
      super(message)
      @status_code = status_code
      @body = body
    end
  end

  class AuthenticationError < ApiError; end
  class PlanNotAllowedError < ApiError; end
  class InsufficientCreditsError < ApiError; end
  class JobConflictError < ApiError; end
  class NotFoundError < ApiError; end
  class BadRequestError < ApiError; end
  class TimeoutError < ApiError; end
end
