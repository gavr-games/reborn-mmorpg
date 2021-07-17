# frozen_string_literal: true

require 'dry-validation'

module Players
  module Validation
    class LoginContract < Dry::Validation::Contract
      params do
        required(:username).filled(:string)
        required(:password).filled(:string)
      end
    end
  end
end
