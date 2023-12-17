require 'minitest/autorun'
require_relative '../main/feature/handling/handle_rc'
require_relative '../main/model/rice_cooker'

class Test < Minitest::Test
    RcHandler.add(1, true)
    RcHandler.add(2, false)
    def test_add_rice_cooker
        assert(2 == RcHandler.rc_list.size)
    end

    def test_change_state
        RcHandler.change_state(2, "is_operational", true)
        RcHandler.change_state(2, "is_cooking", true)
        assert(RcHandler.get_rc(2).is_operational == true)
        assert(RcHandler.get_rc(2).is_cooking == true)

        RcHandler.change_state(1, "is_plugged", true)
        assert(RcHandler.get_rc(1).is_plugged == true)
    end
end