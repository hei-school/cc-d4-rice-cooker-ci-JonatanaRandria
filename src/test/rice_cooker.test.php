<?php

use PHPUnit\Framework\TestCase;

require 'vendor/autoload.php'; // Make sure to include PHPUnit autoload file
require_once '../main/service/Exception/handle_rc.php';
require_once '../main/model/rice_cooker.php';

class Test extends TestCase
{
    public function testAddRiceCooker()
    {
        RcHandler::add(1, true);
        RcHandler::add(2, false);
        $this->assertSame(2, count(RcHandler::rc_list()));
    }

    public function testChangeState()
    {
        RcHandler::change_state(2, "is_free", true);
        RcHandler::change_state(2, "is_cooking", true);
        $this->assertTrue(RcHandler::get_rc(2)->getIsFree());
        $this->assertTrue(RcHandler::get_rc(2)->getIsCooking());

        RcHandler::change_state(1, "is_plugged", true);
        $this->assertTrue(RcHandler::get_rc(1)->getIsPlugged());
    }
}

?>
