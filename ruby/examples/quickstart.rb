# frozen_string_literal: true

$LOAD_PATH.unshift File.expand_path("../lib", __dir__)
require "gmaps_scraper"

client = GmapsScraper::Client.new
me = client.me
puts "plan=#{me['plan']} creditsRemaining=#{me['creditsRemaining']}"

rows = client.scrape("dentists in Austin TX")
puts "scraped #{rows.length} places"
unless rows.empty?
  r = rows.first
  puts "#{r['Name']} #{r['Phone']} #{r['Website']} #{r['Emails']}"
end
