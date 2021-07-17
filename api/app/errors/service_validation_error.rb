class ServiceValidationError < StandardError
  def initialize(validation, model = nil)
      @validation = validation
      @model = model
  end

  def message
      @validation.errors.full_messages
  end

  def validation
      @validation
  end

  def model
      @model
  end
end
