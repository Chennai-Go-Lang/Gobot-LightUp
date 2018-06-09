#!/usr/bin/env ruby

RASPI_HOST = ENV['RASPI_HOST']

require 'http'
device_response = HTTP.get("http://#{RASPI_HOST}:3000/api/robots/Oom/devices").parse

# require 'json'
# device_response = JSON.parse(File.read('/tmp/device-response.json'))

device_ids = device_response["devices"].map {|device| device["name"]}
puts device_ids.join(", ")

COMMANDS = ["On", "Off", "Toggle"]

def get_led_url(led, command)
  "http://#{RASPI_HOST}:3000/api/robots/Oom/devices/#{led}/commands/#{command}"
end

magic_string = "0000000011111111--------"

magic = magic_string.chars.to_a.combination(8).map(&:join).shuffle

magic.each do |design|
  design.chars.to_a.each_with_index.map do |char, index|
    if char == "-"
      command_index = 2
    else
      command_index = char.to_i
    end

    url = get_led_url(device_ids[index], COMMANDS[command_index])

    puts url
    HTTP.get(url)
  end
  sleep 0.5
end
