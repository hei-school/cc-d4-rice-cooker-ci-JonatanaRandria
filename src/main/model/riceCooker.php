<?php

class RiceCooker {
    private $id;
    private $is_free;
    private $is_plugged = false;
    private $is_cooking = false;

    public function __construct($id, $is_free) {
        $this->id = $id;
        $this->is_free = $is_free;
    }

    public function getId() {
        return $this->id;
    }

    public function getIsFree() {
        return $this->is_free;
    }

    public function setIsFree($is_free) {
        $this->is_free = $is_free;
    }

    public function getIsPlugged() {
        return $this->is_plugged;
    }

    public function setIsPlugged($is_plugged) {
        $this->is_plugged = $is_plugged;
    }

    public function getIsCooking() {
        return $this->is_cooking;
    }

    public function setIsCooking($is_cooking) {
        $this->is_cooking = $is_cooking;
    }

    public function __toString() {
        return "Rice cooker : {
            id: {$this->id},
            is_free: {$this->is_free},
            is_cooking: {$this->is_cooking},
            is_plugged: {$this->is_plugged}
        }";
    }
}

?>
