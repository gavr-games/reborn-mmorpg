# frozen_string_literal: true

require 'bcrypt'
require 'jwt'
require_relative '../base_service'
require_relative 'validation/login_contract'

module Players
  class Login < BaseService
    def initialize(params)
      @params = params
    end

    def call
      validate
      find_player
      check_password
      generate_token
      to_hash
    end

    private

    def validate
      contract = Validation::LoginContract.new
      result = contract.call(@params)
      raise ServiceError, error_messages(result) if result.failure?
    end

    def find_player
      @player = Player[username: @params[:username]]
      raise ServiceError, 'Username or password are wrong' if @player.nil?
    end

    def check_password
      password = BCrypt::Password.new(@player.password)
      raise ServiceError, 'Username or password are wrong' if password != @params[:password]
    end

    def generate_token
      @token = JWT.encode(payload(@player), ENV['JWT_SECRET'], 'HS256')
    end

    def to_hash
      {
        token: @token
      }
    end

    def payload(player)
      {
        exp: Time.now.to_i + 60 * 60,
        iat: Time.now.to_i,
        id: player.id
      }
    end
  end
end
