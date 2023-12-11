require_relative "riceCooker"

def main
  puts "Is the electricity cut off? (1 for Yes, 2 for No)"
  puts "1- Yes\n2- No"
  electricity = gets.chomp

  if electricity == "1"
    puts "Mandreta afo na gaz fa tsy poinsa leh rice cooker vo tsisy jiro."
  elsif electricity == "2"
    rice_cooker = RiceCooker.new
    user_prompt(rice_cooker)
  else
    puts "Choose a valid choice."
  end
end

def user_prompt(rice_cooker)
  puts "RICE COOKER:"
  puts "1- Plug in\n2- Unplug\n3- Turn on\n4- Turn off\n5- Add rice\n6- Add water\n7- Show status\n8- Quit"
  user_choice = gets.chomp

  case user_choice
  when "1"
    rice_cooker.plug
  when "2"
    rice_cooker.unplug
  when "3"
    rice_cooker.turn_on
  when "4"
    rice_cooker.turn_off
  when "5"
    rice_cooker.add_rice
  when "6"
    rice_cooker.add_water
  when "7"
    rice_cooker.get_status
  when "8"
    return
  else
    puts "Invalid choice. Please enter a number between 1 and 8."
  end

  # If the user did not choose to quit, call user_prompt recursively
  user_prompt(rice_cooker) unless user_choice == "8"
end

  
  main
  