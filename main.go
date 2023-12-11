package main

import (
	"fmt"

	"rice.cooker/main/module"
)

type RiceCooker struct {
	isPluged  bool
	isEmpty   bool
	isCooking bool
}

func main() {
	var i int = 0
	var choice string

	var riceCooker RiceCooker

	riceCooker.isCooking = false
	riceCooker.isEmpty = true
	riceCooker.isPluged = false

	for {
		fmt.Println("\nRICE COOKER")
		fmt.Println("------------------------")
    fmt.Println("\nMenu:\n")
    fmt.Println("1- Plug In")
    fmt.Println("2- Put meal")
    fmt.Println("3- Power On")
    fmt.Println("4- Power Off")
    fmt.Println("5- Empty")
    fmt.Println("6- Unplug")
    fmt.Println("7- Exit\n")
		fmt.Println("I want to: ")
		fmt.Scan(&choice)

		switch choice {
			case "1":
				res, err := module.PlugIn(riceCooker.isPluged)
				riceCooker.isPluged = res
				if err != nil {
					fmt.Println("Error: ", err)
				} else {
					fmt.Println("\nRice Cooker pluged !\n")
				}
				
			case "2":
				res, err := module.PutSomething(riceCooker.isEmpty)
				riceCooker.isEmpty = res
				if err != nil {
					fmt.Println("Error: ", err)
				} else {
					fmt.Println("\nRice Cooker ready !\n")
				}

			case "3":
				res, err := module.SwitchOn(riceCooker.isCooking, riceCooker.isEmpty, riceCooker.isPluged)
				riceCooker.isCooking = res
				if err != nil {
					fmt.Println("Error: ", err)
				} else {
					fmt.Println("\nRice Cooker On !\n")
				}

			case "4":
				res, err := module.SwitchOff(riceCooker.isCooking)
				riceCooker.isCooking = res
				if err != nil {
					fmt.Println("Error: ", err)
				} else {
					fmt.Println("\nRice Cooker Off !\n")
				}

			case "5":
				res, err := module.Empty(riceCooker.isEmpty, riceCooker.isCooking)
				riceCooker.isEmpty = res
				if err != nil {
					fmt.Println("Error: ", err)
				} else {
					fmt.Println("\nYour rice cooker is empty,put something to cook !\n")
				}
				
			case "6":
				res, err := module.Unplug(riceCooker.isPluged, riceCooker.isCooking)
				riceCooker.isPluged = res
				if err != nil {
					fmt.Println("Error: ", err)
				} else {
					fmt.Println("\nRice Cooker unpluged !\n")
				}


			case "7":
				fmt.Println("Bye !")
				i = 1
		
			default:
				fmt.Println("")
				
			}

		if i == 1 {
			break
		}		
	}
}