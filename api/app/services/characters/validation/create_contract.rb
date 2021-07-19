# frozen_string_literal: true

require 'dry-validation'

module Characters
  module Validation
    class CreateContract < Dry::Validation::Contract
      params do
        required(:name).filled(:string)
        required(:gender).filled(:string)
        required(:player_id).filled(:integer)
      end

      rule(:name) do
        key.failure('must be at least 6 characters') if value.size < 6
        key.failure('already exists') if Character[name: value]&.exists?
      end

      rule(:gender) do
        key.failure('forbidden gender') if !%w{male female}.include?(value)
      end
    end
  end
end
