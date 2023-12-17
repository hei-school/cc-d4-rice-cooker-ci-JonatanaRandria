<?php

require_once '../../model/rice_cooker.php';

class RcHandler {
    private static $rc_list = [];

    public static function add($id, $is_free) {
        $new_rc = new RiceCooker($id, $is_free);
        self::$rc_list[] = $new_rc;
        echo "New rice cooker with id: {$id} successfully added." . PHP_EOL;
    }

    public static function change_state($id, $target_attr, $state) {
        $target_rc = self::get_rc($id);
        if ($target_rc) {
            switch ($target_attr) {
                case 'is_free':
                    $target_rc->setIsFree($state);
                    break;
                case 'is_cooking':
                    $target_rc->setIsCooking($state);
                    break;
                case 'is_plugged':
                    $target_rc->setIsPlugged($state);
                    break;
                default:
                    echo 'The targeted attribute is not valid' . PHP_EOL;
            }
        } else {
            echo "The rice cooker with id: {$id} doesn't exist." . PHP_EOL;
        }
    }

    public static function rc_list() {
        return self::$rc_list;
    }

    public static function get_rc($id) {
        $found = array_filter(self::$rc_list, function ($rc) use ($id) {
            return $rc->getId() == $id;
        });

        if (!empty($found)) {
            return reset($found);
        } else {
            echo "The rice cooker with id: {$id} doesn't exist." . PHP_EOL;
        }
    }
}

?>
