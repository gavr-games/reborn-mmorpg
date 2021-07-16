# frozen_string_literal: true

require 'sinatra'
require_relative 'config/db'
require_relative 'config/loaders'
require_dir 'app'

set :bind, '0.0.0.0'

namespace '/api/v1' do
  post '/players' do
    Players::Register.call(params)
  rescue ServiceValidationError => e
    
  end
end
