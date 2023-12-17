require_relative '../../model/rice_cooker'

module RcHandler
  @rc_list = []

  def self.add(id, is_free)
    new_rc = RiceCooker.new(id, is_free)
    @rc_list.push(new_rc)
    puts "New rice cooker with id: #{id} successfully added."
  end

  def self.change_state(id, target_attr, state)
    target_rc = @rc_list.find { |rc| rc.id == id }
    if target_rc
      case target_attr
      when 'is_free'
        target_rc.is_free = state
      when 'is_cooking'
        target_rc.is_cooking = state
      when 'is_plugged'
        target_rc.is_plugged = state
      else
        puts 'The targeted attribute is not valid'
      end
    else
      puts "The rice cooker with id: #{id} doesn't exist."
    end
  end

  def self.rc_list
    @rc_list
  end

  def self.get_rc(id)
    found = @rc_list.find { |rc| rc.id == id }
    return found if found

    puts "The rice cooker with id: #{id} doesn't exist."
  end
end
