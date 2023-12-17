class RiceCooker
    def initialize(id, is_free)
      @id = id
      @is_free = is_free
      @is_plugged = false
      @is_cooking = false
    end
  
    attr_accessor :is_free, :is_plugged, :is_cooking
    attr_reader :id
  
    def to_s
      "Rice cooker : {
              id: #{id},
              is_free: #{is_free},
              is_cooking: #{is_cooking},
              is_plugged: #{is_plugged}
          }"
    end
  end
  