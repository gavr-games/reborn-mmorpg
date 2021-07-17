# frozen_string_literal: true

require 'dry-validation'

module Players
  module Validation
    class CreateContract < Dry::Validation::Contract
      params do
        required(:email).filled(:string)
        required(:username).filled(:string)
        required(:password).filled(:string)
      end

      rule(:email) do
        unless /\A[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z/i.match?(value)
          key.failure('has invalid format')
        end
        key.failure('already exists') if Player[email: value]&.exists?
      end

      rule(:username) do
        key.failure('must be at least 6 characters') if value.size < 6
        key.failure('already exists') if Player[username: value]&.exists?
      end

      rule(:password) do
        key.failure('must be at least 6 characters') if value.size < 6
      end
    end
  end
end
