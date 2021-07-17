# frozen_string_literal: true

require 'sinatra'
require 'sinatra/contrib'
require 'rack'
require 'rack/contrib'
require 'pry'
require_relative 'config/db'
require_relative 'config/loaders'
require_dir 'app'

set :bind, '0.0.0.0'

use Rack::JSONBodyParser

namespace '/api/v1' do
  post '/players' do
    json Players::Register.call(params.to_h.transform_keys(&:to_sym)).to_hash
  rescue ServiceError => e
    halt(422, json(error: e.message))
  end

  post '/login' do
    json Players::Login.call(params.to_h.transform_keys(&:to_sym)).to_hash
  rescue ServiceError => e
    halt(422, json(error: e.message))
  end
end
