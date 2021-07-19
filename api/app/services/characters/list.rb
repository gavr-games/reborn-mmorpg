# frozen_string_literal: true

require_relative '../base_service'
require_relative 'validation/create_contract'

module Characters
  class List < BaseService
    LIMIT = 4

    def initialize(player)
      @player = player
    end

    def call
      get_list
      to_hash
    end

    private

    def get_list
      @chars = @player.characters
    end

    def to_hash
      @chars.map do |char|
        {
          id: char.id,
          name: char.name,
          gender: char.gender
        }
      end
    end
  end
end
