# frozen_string_literal: true

module Auth
  def protected!
    return if authorized?
    halt(403, json(error: 'Not authorized'))
  end

  def authorized?
    @token = env.fetch('HTTP_AUTHORIZATION', '').slice(7..-1)
    return false if @token.nil?

    begin
      @payload, @header = JWT.decode(@token, ENV['JWT_SECRET'], true, { algorithm: 'HS256'} )

      @exp = @payload['exp']
      if @exp.nil?
        return false
      end

      @exp = Time.at(@exp.to_i)

      if Time.now > @exp
        return false
      end

      @player = Player[@payload['id']]

      return false if @player.nil?

      true
    rescue JWT::DecodeError
      return false
    end
  end
end
