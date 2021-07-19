# frozen_string_literal: true

require_relative '../base_service'
require_relative 'validation/create_contract'

module Characters
  class Get < BaseService
    def initialize(id, player)
      @id = id
      @player = player
    end

    def call
      get_character
      check_player
      to_hash
    end

    private

    def get_character
      @character = Character[@id]
      raise ServiceError, 'character not found' if @character.nil?
    end

    def check_player
      raise ServiceError, 'character not found' if @character.player_id != @player.id
    end

    def to_hash
      {
        id: @character.id,
        name: @character.name,
        gender: @character.gender
      }
    end
  end
end
