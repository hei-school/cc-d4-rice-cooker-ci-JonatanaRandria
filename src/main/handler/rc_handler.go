package main

import (
	"fmt"
)

// RiceCooker represents a rice cooker.
type RiceCooker struct {
	id        int
	isFree    bool
	isPlugged bool
	isCooking bool
}

// NewRiceCooker creates a new RiceCooker instance.
func NewRiceCooker(id int, isFree bool) *RiceCooker {
	return &RiceCooker{
		id:        id,
		isFree:    isFree,
		isPlugged: false,
		isCooking: false,
	}
}

// RcHandler manages rice cookers.
type RcHandler struct {
	rcList []*RiceCooker
}

// Add adds a new rice cooker to the list.
func (rh *RcHandler) Add(id int, isFree bool) {
	newRC := NewRiceCooker(id, isFree)
	rh.rcList = append(rh.rcList, newRC)
	fmt.Printf("New rice cooker with id: %d successfully added.\n", id)
}

// ChangeState changes the state of a rice cooker attribute.
func (rh *RcHandler) ChangeState(id int, targetAttr string, state bool) {
	targetRC := rh.getRc(id)
	if targetRC != nil {
		switch targetAttr {
		case "isFree":
			targetRC.isFree = state
		case "isCooking":
			targetRC.isCooking = state
		case "isPlugged":
			targetRC.isPlugged = state
		default:
			fmt.Println("The targeted attribute is not valid")
		}
	} else {
		fmt.Printf("The rice cooker with id: %d doesn't exist.\n", id)
	}
}

// RcList returns the list of rice cookers.
func (rh *RcHandler) RcList() []*RiceCooker {
	return rh.rcList
}

// GetRc returns a rice cooker by ID.
func (rh *RcHandler) GetRc(id int) *RiceCooker {
	found := rh.getRc(id)
	if found != nil {
		return found
	}
	fmt.Printf("The rice cooker with id: %d doesn't exist.\n", id)
	return nil
}

// getRc finds a rice cooker by ID.
func (rh *RcHandler) getRc(id int) *RiceCooker {
	for _, rc := range rh.rcList {
		if rc.id == id {
			return rc
		}
	}
	return nil
}

