package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var rcHandler RcHandler

// RcHandler struct for handling rice cookers
type RcHandler struct {
	rcList []*RiceCooker
}

// RiceCooker struct representing a rice cooker
type RiceCooker struct {
	id          int
	isFree      bool
	isPlugged   bool
	isCooking   bool
}

func newRiceCooker(id int, isFree bool) *RiceCooker {
	return &RiceCooker{id: id, isFree: isFree}
}

func (rc *RiceCooker) String() string {
	return fmt.Sprintf("Rice cooker: {id: %d, isFree: %t, isCooking: %t, isPlugged: %t}", rc.id, rc.isFree, rc.isCooking, rc.isPlugged)
}

func (handler *RcHandler) add(id int, isFree bool) {
	newRc := newRiceCooker(id, isFree)
	handler.rcList = append(handler.rcList, newRc)
	fmt.Printf("New rice cooker with id: %d successfully added.\n", id)
}

func (handler *RcHandler) changeState(id int, targetAttr string, state bool) {
	targetRc := handler.getRc(id)
	if targetRc != nil {
		switch targetAttr {
		case "isFree":
			targetRc.isFree = state
		case "isCooking":
			targetRc.isCooking = state
		case "isPlugged":
			targetRc.isPlugged = state
		default:
			fmt.Println("The targeted attribute is not valid")
		}
	} else {
		fmt.Printf("The rice cooker with id: %d doesn't exist.\n", id)
	}
}

func (handler *RcHandler) rcListString() string {
	result := "Rice cookers:\n"
	for _, rc := range handler.rcList {
		result += rc.String() + "\n"
	}
	return result
}

func (handler *RcHandler) getRc(id int) *RiceCooker {
	for _, rc := range handler.rcList {
		if rc.id == id {
			return rc
		}
	}
	fmt.Printf("The rice cooker with id: %d doesn't exist.\n", id)
	return nil
}

func showMenu(message string, valid []int, release int, operation func(int)) {
	fmt.Print(message)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	choice, err := strconv.Atoi(input.Text())
	if err == nil && contains(valid, choice) {
		operation(choice)

		if choice != release {
			showMenu(message, valid, release, operation)
		}
	} else {
		fmt.Printf("Invalid, your input must be one of %v\n", valid)
		showMenu(message, valid, release, operation)
	}
}

func contains(slice []int, elem int) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

func main() {
	rcHandler = RcHandler{}

	showMenu("Choose an action: ", []int{1, 2, 3, 4}, 4, func(choice int) {
		switch choice {
		case 1:
			addAction()
		case 2:
			handleRCMenu()
		case 3:
			cookMenu()
		case 4:
			fmt.Println("Goodbye, enjoy your meal!")
		}
	})
}

func getRcUserInput() (int, bool) {
	fmt.Print("Enter your rice cooker id (Must be a number): ")
	var id int
	fmt.Scanln(&id)

	fmt.Print("Is this rice cooker free? (1. yes, 2. no): ")
	var isOperational int
	fmt.Scanln(&isOperational)

	return id, isOperational == 1
}

func addAction() {
	id, isFree := getRcUserInput()

	if id > 0 {
		rcHandler.add(id, isFree)
	} else {
		fmt.Println("Invalid choice, the id must be a positive number.")
		addAction()
	}
}

func handleRCMenu() {
	menuMessage := "Choose an action:\n" +
		"1. List rice cookers\n" +
		"2. Plug in/out rice cooker\n" +
		"3. Change free state\n" +
		"4. Return to main menu\n" +
		"I wanna: "

	showMenu(menuMessage, []int{1, 2, 3, 4}, 4, func(choice int) {
		switch choice {
		case 1:
			fmt.Println(rcHandler.rcListString())
		case 2:
			updateRCState(choice, "isPlugged", "is now plugged out", "is now plugged in", "updated")
		case 3:
			updateRCState(choice, "isFree", "is now non-free", "is now free", "updated")
		case 4:
			fmt.Println("Returning to the main menu.")
		}
	})
}

func updateRCState(choice int, attribute string, oppositeValueStr string, oppositeMessage string, currentMessage string) {
	fmt.Print("Enter the rice cooker id (Must be a number): ")
	var id int
	fmt.Scanln(&id)

	target := rcHandler.getRc(id)
	if target != nil {
		currentState := false
		switch attribute {
		case "isPlugged":
			currentState = target.isPlugged
		case "isFree":
			currentState = target.isFree
		}

		oppositeValue := !currentState
		rcHandler.changeState(id, attribute, oppositeValue)
		fmt.Printf("Rice cooker id: %d %s\n", id, mapBoolToStr(currentState, oppositeMessage, currentMessage))
	}
}

func mapBoolToStr(b bool, trueStr string, falseStr string) string {
	if b {
		return trueStr
	}
	return falseStr
}

func cookMenu() {
	menuMessage := "Choose an action:\n" +
		"1. Cook rice\n" +
		"2. Boil water\n" +
		"3. Return to main menu\n" +
		"Your choice: "

	showMenu(menuMessage, []int{1, 2, 3}, 3, func(choice int) {
		switch choice {
		case 1:
			cookRice()
		case 2:
			boilWater()
		case 3:
			fmt.Println("Returning to the main menu.")
		}
	})
}

func getBoilWaterUserInput() (int, int) {
	fmt.Print("Enter the rice cooker ref number to use: ")
	var rc int
	fmt.Scanln(&rc)

	fmt.Print("Enter the number of water cups: ")
	var cups int
	fmt.Scanln(&cups)

	return rc, cups
}

func boilWater() {
	rc, cups := getBoilWaterUserInput()

	if rc > 0 && cups > 0 {
		if cups <= 0 {
			fmt.Println("Add more water")
			cookRice()
		} else if target := rcHandler.getRc(rc); target != nil {
			startCooking(rc)
		}
	} else {
		fmt.Println("Invalid input. Both rice cooker id and water cups must be numbers.")
	}
}

func getCookRiceUserInput() (int, int, int) {
	fmt.Print("Enter the rice cooker id to use: ")
	var rcID int
	fmt.Scanln(&rcID)

	fmt.Print("Enter the number of water cups: ")
	var cups int
	fmt.Scanln(&cups)

	fmt.Print("Enter the number of rice cups (should be 1/2 of water cups number): ")
	var rice int
	fmt.Scanln(&rice)

	return rcID, cups, rice
}

func cookRice() {
	rcID, cups, rice := getCookRiceUserInput()

	if rcID > 0 && cups > 0 && rice > 0 && cups >= rice*2 {
		startCooking(rcID)
	} else {
		fmt.Println("Invalid input. Check your rice cooker id, water cups, and rice cups.")
	}
}

func startCooking(id int) {
	target := rcHandler.getRc(id)
	if target != nil {
		rcHandler.changeState(id, "isCooking", true)
		fmt.Printf("Rice cooker id: %d started cooking\n", id)
	}
}
