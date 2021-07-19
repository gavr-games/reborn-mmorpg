# frozen_string_literal: true

class Character < Sequel::Model(DB)
  many_to_one :player
end
