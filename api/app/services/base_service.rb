# frozen_string_literal: true

class BaseService
  def self.call(*args, &block)
    new(*args, &block).call
  end

  private

  # Accepts dry validation result
  def error_messages(result)
    result.errors(full: true).map(&:text).join(', ')
  end
end
