# frozen_string_literal: true

class CreateCharactersTable < Sequel::Migration
  def up
    create_table :characters do
      primary_key :id
      foreign_key :player_id, :players
      String :name, null: false
      String :gender, null: false
      DateTime :created_at, null: false
      DateTime :updated_at, null: false
    end
  end

  def down
    drop_table :characters
  end
end
