# frozen_string_literal: true

require 'sinatra'

set :bind, '0.0.0.0'

get '/api' do
  'Hello world!'
end
