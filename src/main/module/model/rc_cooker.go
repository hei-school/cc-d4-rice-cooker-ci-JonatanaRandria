package main

import "fmt"

type RiceCooker struct {
	id         int
	isFree     bool
	isPlugged  bool
	isCooking  bool
}

func NewRiceCooker(id int, isFree bool) *RiceCooker {
	return &RiceCooker{
		id:        id,
		isFree:    isFree,
		isPlugged: false,
		isCooking: false,
	}
}

func (rc *RiceCooker) String() string {
	return fmt.Sprintf("Rice cooker : { id: %d, isFree: %t, isCooking: %t, isPlugged: %t }", rc.id, rc.isFree, rc.isCooking, rc.isPlugged)
}

