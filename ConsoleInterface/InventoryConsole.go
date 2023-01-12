package ConsoleInterface

import (
	"fmt"

	"github.com/1234bharathi/GOLANGASSIGN/Support"

	Handlers "github.com/1234bharathi/GOLANGASSIGN/Handlers/Inventory"
)

func InventoryConsole() {
	fmt.Println("Hi, you are here to perform CRUD operations on 'product' table")
	fmt.Println("Please choose the task to perform")
	fmt.Printf("1.Insert\n2.Delete\n3.Get All Inventory  \n4.Update\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		InsertInventory()
	} else if choice == 2 {
		DeleteInventory()
	} else if choice == 3 {
		GetInventory()
	} else if choice == 4 {
		UpdateInventory()
	}
	return
}

func DeleteInventory() {
	Handlers.DeleteInventoryHandlerConsole()
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		InventoryConsole()
	}
	return

}
func InsertInventory() {
	Inventory := Support.GetInventoryInput()

	Handlers.AddInventoryConsoleHandler(Inventory)
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		InventoryConsole()
	} else {
		return
	}
}

func GetInventory() {
	Handlers.GetAllInventoryConsoleHandler()

	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		InventoryConsole()
	}
	return

}

// func GetAllInventory() {
// 	Handlers.GetAllInventoryHandlerConsole()
// 	var cont string
// 	_, err := fmt.Scanln(&cont)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	if cont == "yes" {
// 		InventoryConsole()
// 	}
// 	return
// }

func UpdateInventory() {
	Handlers.UpdateInventoryConsole()
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		InventoryConsole()
	} else {
		return
	}
}
