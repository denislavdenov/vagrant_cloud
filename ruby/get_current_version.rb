#!/usr/bin/env ruby

puts "The box of which user do you want?"
var1 = gets
puts "The url of which box of this user do you want?"
var2 = gets

require "http"
require "json"
require "rubygems"

api = HTTP.persistent("https://app.vagrantup.com").headers(
  "Authorization" => "Bearer #{ENV['VAGRANT_CLOUD_TOKEN']}"
)

response = api.get("/api/v1/search", params: {
  q: "#{var1}/#{var2}",
  provider: "vmware_desktop"
})

if response.status.success?
  # Success, the response attributes are available here.
	json=response.parse["boxes"][0]["current_version"]["providers"]
		json.each do |item|
	          item.each do |key,value|
			p value if key == "download_url"
		  end
		end
else
  # Error, inspect the `errors` key for more information.
  p response.code, response.body
end


