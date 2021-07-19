# frozen_string_literal: true

require_relative '../base_service'
require_relative 'validation/create_contract'

module Characters
  class Create < BaseService
    LIMIT = 4

    def initialize(params)
      @params = params
    end

    def call
      validate_params
      validate_limit
      create
    end

    private

    def validate_params
      contract = Validation::CreateContract.new
      result = contract.call(@params)
      raise ServiceError, error_messages(result) if result.failure?
    end

    def validate_limit
      characters_count = Character.count(player_id: @params[:player_id])
      raise ServiceError, "Cannot create more than #{LIMIT} characters" if characters_count >= LIMIT
    end

    def create
      Character.create(
        name: @params[:name],
        gender: @params[:gender],
        player_id: @params[:player_id]
      )
    end
  end
end
