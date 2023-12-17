def show_menu(message, valid, release, operation)
    print message
    input = gets.chomp
    if input.match?(/^\d+$/)
      choice = input.to_i
      if valid.include?(choice)
        operation.call(choice)
        if choice != release
          show_menu(message, valid, release,
                    operation)
        end
        nil if choice == release
      else
        puts "You choose the wrong number,please take one of #{valid}"
        show_menu(message, valid, release, operation)
      end
    else
      puts "Invalid, your input must be one of #{valid}"
      show_menu(message, valid, release, operation)
    end
  end
  