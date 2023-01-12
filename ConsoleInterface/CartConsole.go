package ConsoleInterface

import (
	"fmt"

	Handlers "github.com/1234bharathi/GOLANGASSIGN/Handlers/Cart"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func CartConsole() {
	fmt.Println("Hi, you are here to perform CRUD operations on 'product' table")
	fmt.Println("Please choose the task to perform")
	fmt.Printf("1.Create Cart\n2.Add Item to Cart\n3.Add Multiple Items to Cart\n4.Delete Cart \n5.Total price of cart\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		CreateCart()
	} else if choice == 2 {
		AddItemToCart()
	} else if choice == 3 {
		AddMultipleItemsToCart()
	} else if choice == 4 {
		DeleteCategory()
	}
	//  else if choice == 4 {
	// 	GetAllCategory()
	// } else if choice == 5 {
	// 	UpdateCategory()
	// }
	return
}

func CreateCart() {

	Handlers.CreateCartConsoleHandler()
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		CartConsole()
	} else {
		return
	}
}
func AddMultipleItemsToCart() {
	Handlers.AddItemsToCartConsoleHandler()
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		CartConsole()
	} else {
		return
	}
}

func AddItemToCart() {
	NewCart := Support.GetCartInput()
	Handlers.AddToCartConsoleHandler(NewCart)
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		CartConsole()
	} else {
		return
	}
}

// func GetCategory() {
// 	Handlers.GetCategoryConsoleHandler()

// 	fmt.Println("Do you want to continue? (yes or no)")
// 	var cont string
// 	_, err := fmt.Scanln(&cont)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	if cont == "yes" {
// 		CategoryConsole()
// 	}
// 	return

// }

// func GetAllCategory() {
// 	Handlers.GetAllCategoryHandlerConsole()
// 	var cont string
// 	_, err := fmt.Scanln(&cont)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	if cont == "yes" {
// 		CategoryConsole()
// 	}
// 	return
// }

// func UpdateCategory() {
// 	Handlers.UpdateCategoryConsole()
// 	fmt.Println("Do you want to continue? (yes or no)")
// 	var cont string
// 	_, err := fmt.Scanln(&cont)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	if cont == "yes" {
// 		CategoryConsole()
// 	} else {
// 		return
// 	}
// }

func DeleteCart() {
	Handlers.DeleteCartConsoleHandler()
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		CartConsole()
	}
	return

}
