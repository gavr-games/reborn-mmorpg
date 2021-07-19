# frozen_string_literal: true

class Player < Sequel::Model(DB)
  one_to_many :characters
end
