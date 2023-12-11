class RiceCooker
  def initialize
    @is_plugged = false
    @power_on = false
    @rice_inserted = false
    @water_inserted = false
    @rice_quantity = 0
    @water_quantity = 0
    @capacity = 900
  end

  def plug
    if @is_plugged
      puts "Rice Cooker is already plugged in."
    else
      @is_plugged = true
      puts "Rice Cooker plugged in."
    end
  end

  def unplug
    if !@is_plugged
      puts "The rice cooker is already unplugged."
    else
      @is_plugged = false
      puts "The rice cooker has been unplugged."
    end
  end

  def turn_on
    if @power_on
      puts "The Rice Cooker is already on."
    elsif !@is_plugged
      puts "The rice cooker should be plugged before turning on."
    elsif !@rice_inserted
      puts "You should insert rice."
    elsif !@water_inserted
      puts "You should insert water."
    else
      @power_on = true
      puts "Rice Cooker turned on, start cooking."
    end
  end

  def turn_off
    if !@power_on
      puts "The Rice Cooker is already off."
    else
      @power_on = false
      puts "Rice Cooker turned off."
    end
  end

  def add_rice
    if @rice_inserted
      puts "Rice is already present."
    else
      print "Rice quantity: "
      quantity_str = gets.chomp
      quantity = quantity_str.to_i

      if @water_inserted
        if quantity <= @water_quantity && quantity + @rice_quantity <= @capacity
          @rice_quantity = quantity
          @rice_inserted = true
          puts "Rice has been added."
        else
          puts "Not enough space or insufficient water quantity."
        end
      else
        if quantity <= @capacity
          @rice_quantity = quantity
          @rice_inserted = true
          puts "Rice has been added."
        else
          puts "Not enough space."
        end
      end
    end
  end

  def add_water
    if @water_inserted
      puts "Water is already present."
    else
      print "Water quantity: "
      quantity_str = gets.chomp
      quantity = quantity_str.to_i

      if @rice_inserted
        if quantity <= @capacity - @rice_quantity && quantity <= @capacity
          @water_quantity = quantity
          @water_inserted = true
          puts "Water has been added."
        else
          puts "Not enough space or insufficient capacity for water with current rice quantity."
        end
      else
        if quantity <= @capacity
          @water_quantity = quantity
          @water_inserted = true
          puts "Water has been added."
        else
          puts "Not enough space."
        end
      end
    end
  end

  def get_status
    if @is_plugged
      puts "Plugged."
    else
      puts "Unplugged."
    end

    if @power_on
      puts "ON."
    else
      puts "OFF."
    end

    if @rice_inserted
      puts "Rice inserted: #{@rice_quantity}"
    end

    if @water_inserted
      puts "Water inserted: #{@water_quantity}"
    end
  end
end
