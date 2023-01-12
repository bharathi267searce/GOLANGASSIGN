package application

import (
	"fmt"

	"github.com/1234bharathi/GOLANGASSIGN/ConsoleInterface"
)

func Console() {
	fmt.Println("WELCOME TO CONSOLE INTERFACE")
	fmt.Println("Select the table to perform the task")
	fmt.Printf("1.Product\n2.Category\n3.Inventory\n4.Cart\n")
	fmt.Println("Enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}

	if choice == 1 {
		ConsoleInterface.ProductConsole()
	} else if choice == 2 {
		ConsoleInterface.CategoryConsole()
	} else if choice == 3 {
		ConsoleInterface.InventoryConsole()
	} else if choice == 4 {
		ConsoleInterface.CartConsole()
	}
	return
}
