# frozen_string_literal: true

require 'sequel'

DB = Sequel.connect(
  adapter: 'postgres',
  host: ENV['POSTGRES_HOST'],
  database: ENV['POSTGRES_DB'],
  user: ENV['POSTGRES_USER'],
  password: ENV['POSTGRES_PASSWORD']
)

Sequel.extension :migration
Sequel::Model.plugin :force_encoding, 'UTF-8'
Sequel::Model.plugin :timestamps, update_on_create: true
Sequel::Migrator.run(DB, File.join(File.dirname(__FILE__), '../migrations'))
