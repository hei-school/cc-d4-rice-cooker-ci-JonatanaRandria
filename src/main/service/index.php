<?php

require_once './utils/rc_utils.php';
require_once './service/Exception.php';

const MAIN_MENU_MESSAGE = "Welcome on Vilany. Choose one action from the list below:
                        1. Add a new rice cooker
                        2. Handle rice cookers
                        3. Cook
                        4. Leave
                    I wanna : ";
const VALID_CHOICES = [1, 2, 3, 4];

function show()
{
    show_menu(MAIN_MENU_MESSAGE, VALID_CHOICES, 4, 'choose_action');
}

function get_user_input()
{
    echo 'Enter your rice cooker id (Must be a number) :' . PHP_EOL;
    $id = trim(fgets(STDIN));

    echo "Is this rice cooker free ?
          1. yes
          2. no
      Your choice: " . PHP_EOL;
    $is_operational = trim(fgets(STDIN));

    return [$id, $is_operational];
}

function add_action()
{
    list($id, $is_operational) = get_user_input();

    if (preg_match('/^\d+$/', $id)) {
        switch ($is_operational) {
            case '1':
                RcHandler::add((int)$id, true);
                break;
            case '2':
                RcHandler::add((int)$id, false);
                break;
            default:
                echo 'Invalid choice' . PHP_EOL;
                add_action();
                break;
        }
    } else {
        echo 'Invalid, the choice must be a number' . PHP_EOL;
        add_action();
    }
}

function handle_rc_menu()
{
    $menu_message = "Choose an action:
                        1. List rice cookers
                        2. Plug in/out rice cooker
                        3. Change free state
                        4. Return to main menu
                    I wanna : ";

    show_menu($menu_message, [1, 2, 3, 4], null, 'handle_rc');
}

function update_rc_state($choice, $attribute, $opposite_value_str, $opposite_message,
                         $current_message)
{
    echo 'Enter the rice cooker id (Must be a number) :' . PHP_EOL;
    $id = trim(fgets(STDIN));

    if (preg_match('/^\d+$/', $id)) {
        $to_update = RcHandler::get_rc((int)$id);
        if ($to_update) {
            $current_state = $to_update->$attribute;
            $opposite_value = $current_state ? false : true;
            RcHandler::change_state((int)$id, $attribute, $opposite_value);
            echo "Rice cooker id: {$id} {$current_state ? $opposite_message : $current_message}" . PHP_EOL;
        }
    } else {
        echo 'Invalid, your choice must be a number' . PHP_EOL;
        $choice();
    }
}

function plug_rc_in_out($choice)
{
    update_rc_state($choice, 'is_plugged', 'is now plugged out', 'is now plugged in', 'updated');
}

function change_operational_state($choice)
{
    update_rc_state($choice, 'is_operational', 'is now non-free', 'is now free', 'updated');
}

function start_cooking($id)
{
    $to_update = RcHandler::get_rc((int)$id);
    if (!$to_update) {
        return;
    }

    RcHandler::change_state((int)$id, 'is_cooking', true);
    echo "Rice cooker id: {$id} started cooking" . PHP_EOL;
}

function handle_rc($choice)
{
    switch ($choice) {
        case 1:
            echo RcHandler::rc_list() . PHP_EOL;
            show();
            break;
        case 2:
            plug_rc_in_out($choice);
            show();
            break;
        case 3:
            change_operational_state($choice);
            show();
            break;
        case 4:
            show();
            break;
    }
}

function cook_menu()
{
    $menu_message = "Choose an action:
                        1. Cook rice
                        2. Boil water
                        3. Return to main menu
                    Your choice: ";

    show_menu($menu_message, [1, 2, 3], null, 'handle_cooks');
}

function handle_cooks($choice)
{
    switch ($choice) {
        case 1:
            cook_rice();
            show();
            break;
        case 2:
            boil_water();
            show();
            break;
        case 3:
            show();
            break;
    }
}

function user_input_for_boil_water()
{
    echo 'Enter the rice cooker ref number to use: ' . PHP_EOL;
    $rc = trim(fgets(STDIN));
    echo 'Enter the number of water cups: ' . PHP_EOL;
    $cups = trim(fgets(STDIN));

    return [$rc, $cups];
}

function boil_water()
{
    list($rc, $cups) = user_input_for_boil_water();

    if (preg_match('/^\d+$/', $rc) && preg_match('/^\d+$/', $cups)) {
        if ($cups <= 0) {
            echo 'Add more water' . PHP_EOL;
            cook_rice();
        } elseif (check_rc((int)$rc)) {
            start_cooking((int)$rc);
        } else {
            show();
        }
    } else {
        echo 'Invalid input. Both rice cooker id and water cups must be numbers.' . PHP_EOL;
    }
}

function cook_rice()
{
    echo 'Enter the rice cooker id to use: ' . PHP_EOL;
    $rc_id = trim(fgets(STDIN));
    echo 'Enter the number of water cups: ' . PHP_EOL;
    $cups = trim(fgets(STDIN));
    echo 'Enter the number of rice cups (should be 1/2 of water cups number): ' . PHP_EOL;
    $rice = trim(fgets(STDIN));

    if (valid_rc($rc_id) && valid_cups($cups) && valid_rice($rice, $cups)) {
        start_cooking((int)$rc_id);
    } else {
        show();
    }
}

function valid_rc($rc_id)
{
    if (preg_match('/^\d+$/', $rc_id)) {
        return true;
    }

    echo 'Invalid rice cooker id.' . PHP_EOL;
    return false;
}

function valid_cups($cups)
{
    if (preg_match('/^\d+$/', $cups) && $cups > 0) {
        return true;
    }

    echo 'Invalid, the water cups must be a positive number.' . PHP_EOL;
    return false;
}

function valid_rice($rice, $cups)
{
    if (preg_match('/^\d+$/', $rice) && $rice > 0 && $cups >= $rice * 2) {
        return true;
    }

    echo 'Invalid, the rice cups must be a positive number,
        and less than or equal to half of the water cups.' . PHP_EOL;
    return false;
}

function check_rc($id)
{
    $target = RcHandler::get_rc($id);
    if (!$target) {
        return false;
    }

    $error_message = '';
    $error_message .= cooking_error_message($id, $target) ?? '';
    $error_message .= operational_error_message($id, $target) ?? '';
    $error_message .= plugged_error_message($id, $target) ?? '';

    echo $error_message;

    return $error_message === '';
}

function cooking_error_message($id, $target)
{
    return $target->is_cooking ? "Rice cooker id:{$id} is still cooking. " : null;
}

function operational_error_message($id, $target)
{
    return !$target->is_operational ? "Rice cooker id:{$id} is not free. " : null;
}

function plugged_error_message($id, $target)
{
    return !$target->is_plugged ? "Rice cooker id:{$id} is not plugged in. " : null;
}

function choose_action($choice)
{
    switch ($choice) {
        case 1:
            add_action();
            break;
        case 2:
            handle_rc_menu();
            break;
        case 3:
            cook_menu();
            break;
        case 4:
            echo 'Goodbye,enjoy your meal !' . PHP_EOL;
            break;
    }
}

show();

?>
