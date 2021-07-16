# frozen_string_literal: true

require_relative '../base_service'

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
    end

    def create
    end
  end
end
