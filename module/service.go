package module

import "errors"

func PlugIn(status bool) (bool, error) {
	if status {
		return true, errors.New("Rice Cooker already pluged")
	} 
	return true, nil
}

func Unplug(isPluged bool, isCooking bool) (bool, error) {
	if !isPluged {
		return true, errors.New("Rice Cooker already unpluged")
	} else if isCooking {
		return true, errors.New("Rice Cooker is already in use")
	}
	return false, nil
}

func SwitchOn(isCooking bool, isEmpty bool, isPluged bool) (bool, error) {
	if isCooking {
		return true, errors.New("Rice Cooker already On")
		} else if isEmpty {
		return false, errors.New("Rice Cooker already Empty")
		} else if !isPluged {
		return false, errors.New("Rice Cooker already Unpluged")
	}
	return true, nil
}

func SwitchOff(isCooking bool) (bool, error) {
	if !isCooking {
		return false, errors.New("Rice Cooker already Off")
	}
	return false, nil
}

func PutSomething(isEmpty bool) (bool, error) {
	if !isEmpty {
		return false, errors.New("Rice cooker is already full")
	}
	return false, nil
}

func Empty(isEmpty bool, isCooking bool) (bool, error) {
	if isEmpty {
		return true, errors.New("Rice cooker is already empty")
	} else if isCooking {
		return false, errors.New("Rice cooker is already working")
	}
	return true, nil
}
