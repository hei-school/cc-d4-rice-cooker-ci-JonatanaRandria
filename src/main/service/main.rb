require_relative './utils/rc_utils'
require_relative './service/Exception'

MAIN_MENU_MESSAGE = "Welcome on Vilany. Choose one action from the list below:
                        1. Add a new rice cooker
                        2. Handle rice cookers
                        3. Cook
                        4. Leave
                    I wanna : "
VALID_CHOICES = [1, 2, 3, 4].freeze
def show
  show_menu(MAIN_MENU_MESSAGE, VALID_CHOICES, 4, method(:choose_action))
end

def get_user_input
  puts 'Enter your rice cooker id (Must be a number) :'
  id = gets.chomp
  puts "Is this rice cooker free ?
          1. yes
          2. no
      Your choice: "
  is_operational = gets.chomp
  [id, is_operational]
end

def add_action(&block)
  id, is_operational = get_user_input

  if id.match?(/^\d+$/)
    case is_operational
    when '1'
      RcHandler.add(id.to_i, true)
    when '2'
      RcHandler.add(id.to_i, false)
    else
      puts 'Invalid choice'
      add_action(&block)
    end
  else
    puts 'Invalid, the choice must be a number'
    add_action(&block)
  end
end

def handle_rc_menu
  menu_message = "Choose an action:
                        1. List rice cookers
                        2. Plug in/out rice cooker
                        3. Change free state
                        4. Return to main menu
                    I wanna : "

  show_menu(menu_message, [1, 2, 3, 4], nil, method(:handle_rc))
end

def update_rc_state(choice, attribute, _opposite_value_str, opposite_message,
                    current_message)
  puts 'Enter the rice cooker id (Must be a number) :'
  id = gets.chomp
  if id.match?(/^\d+$/)
    to_update = RcHandler.get_rc(id.to_i)
    if to_update
      current_state = to_update.public_send(attribute)
      opposite_value = current_state ? false : true
      RcHandler.change_state(id.to_i, attribute, opposite_value)
      puts "Rice cooker id: #{id} #{current_state ? opposite_message : current_message}"
    end
  else
    puts 'Invalid, your choice must be a number'
    send(choice)
  end
end

def plug_rc_in_out(choice)
  update_rc_state(choice, 'is_plugged', 'is now plugged out',
                  'is now plugged in', 'updated')
end

def change_operational_state(choice)
  update_rc_state(choice, 'is_operational', 'is now non-free',
                  'is now free', 'updated')
end

def start_cooking(id)
  to_update = RcHandler.get_rc(id.to_i)
  return unless to_update

  RcHandler.change_state(id.to_i, 'is_cooking', true)
  puts "Rice cooker id: #{id} started cooking"
end

def handle_rc(choice)
  case choice
  when 1
    puts RcHandler.rc_list
    show
  when 2
    plug_rc_in_out(choice)
    show
  when 3
    change_operational_state(choice)
    show
  when 4
    show
  end
end

def cook_menu
  menu_message = "Choose an action:
                        1. Cook rice
                        2. Boil water
                        3. Return to main menu
                    Your choice: "

  show_menu(menu_message, [1, 2, 3], nil, method(:handle_cooks))
end

def handle_cooks(choice)
  case choice
  when 1
    cook_rice
    show
  when 2
    boil_water
    show
  when 3
    show
  end
end

def user_input_for_boil_water
  puts 'Enter the rice cooker ref number to use: '
  rc = gets.chomp
  puts 'Enter the number of water cups: '
  cups = gets.chomp
  [rc, cups]
end

def boil_water
  rc, cups = user_input_for_boil_water

  if rc.match?(/^\d+$/) && cups.match?(/^\d+$/)
    if cups.to_i <= 0
      puts 'Add more water'
      cook_rice
    elsif check_rc(rc.to_i)
      start_cooking(rc.to_i)
    else
      show
    end
  else
    puts 'Invalid input. Both rice cooker id and water cups must be numbers.'
  end
end

def cook_rice
  puts 'Enter the rice cooker id to use: '
  rc_id = gets.chomp
  puts 'Enter the number of water cups: '
  cups = gets.chomp
  puts 'Enter the number of rice cups (should be 1/2 of water cups number): '
  rice = gets.chomp

  if valid_rc?(rc_id) && valid_cups?(cups) && valid_rice?(rice, cups)
    start_cooking(rc_id.to_i)
  else
    show
  end
end

def valid_rc?(rc_id)
  return true if rc_id.match?(/^\d+$/)

  puts 'Invalid rice cooker id.'
  false
end

def valid_cups?(cups)
  return true if cups.match?(/^\d+$/) && cups.to_i.positive?

  puts 'Invalid, the water cups must be a positive number.'
  false
end

def valid_rice?(rice, cups)
  return true if rice.match?(/^\d+$/) && rice.to_i.positive? && cups.to_i >= rice.to_i * 2

  puts 'Invalid, the rice cups must be a positive number,
        and less than or equal to half of the water cups.'
  false
end

def check_rc(id)
  target = RcHandler.get_rc(id)
  return false unless target

  error_message = ''
  error_message += cooking_error_message(id, target) if target.is_cooking == true
  error_message += operational_error_message(id, target) if target.is_operational == false
  error_message += plugged_error_message(id, target) if target.is_plugged == false

  print error_message unless error_message.empty?
  !error_message.empty?
end

def cooking_error_message(id, target)
  "Rice cooker id:#{id} is still cooking. " if target.is_cooking == true
end

def operational_error_message(id, target)
  "Rice cooker id:#{id} is not free. " if target.is_operational == false
end

def plugged_error_message(id, target)
  "Rice cooker id:#{id} is not plugged in. " if target.is_plugged == false
end

def choose_action(choice)
  case choice
  when 1
    add_action
  when 2
    handle_rc_menu
  when 3
    cook_menu
  when 4
    puts 'Goodbye,enjoy your meal !'
    nil
  end
end

show
