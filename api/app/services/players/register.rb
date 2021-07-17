# frozen_string_literal: true

require 'bcrypt'
require_relative '../base_service'
require_relative 'validation/create_contract'

module Players
  class Register < BaseService
    def initialize(params)
      @params = params
    end

    def call
      validate
      create
    end

    private

    def validate
      contract = Validation::CreateContract.new
      result = contract.call(@params)
      raise ServiceError, error_messages(result) if result.failure?
    end

    def create
      Player.create(
        username: @params[:username],
        email: @params[:email],
        password: BCrypt::Password.create(@params[:password])
      )
    end
  end
end
