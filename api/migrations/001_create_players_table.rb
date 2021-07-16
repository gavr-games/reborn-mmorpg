# frozen_string_literal: true

class CreatePlayersTable < Sequel::Migration
  def up
    create_table :players do
      primary_key :id
      String :username, null: false
      String :password, null: false
      String :email, null: false
      DateTime :created_at, null: false
      DateTime :updated_at, null: false
    end
  end

  def down
    drop_table :players
  end
end
