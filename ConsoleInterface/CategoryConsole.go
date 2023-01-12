package ConsoleInterface

import (
	"fmt"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"

	Handlers "github.com/1234bharathi/GOLANGASSIGN/Handlers/Category"
)

func CategoryConsole() {
	fmt.Println("Hi, you are here to perform CRUD operations on 'product' table")
	fmt.Println("Please choose the task to perform")
	fmt.Printf("1.Insert\n2.Read with Category id\n3.Delete\n4.Get All Category by page number  \n5.Update\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		InsertCategory()
	} else if choice == 2 {
		GetCategory()
	} else if choice == 3 {
		DeleteCategory()
	} else if choice == 4 {
		GetAllCategory()
	} else if choice == 5 {
		UpdateCategory()
	}
	return
}

func DeleteCategory() {
	Handlers.DeleteCategoryConsoleHandler()
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		CategoryConsole()
	}
	return

}
func InsertCategory() {
	var NewCategory Datastructures.Category_master = Support.GetCategoryInput()

	Handlers.AddCategoryConsoleHandle(NewCategory)
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		CategoryConsole()
	} else {
		return
	}
}

func GetCategory() {
	Handlers.GetCategoryConsoleHandler()

	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		CategoryConsole()
	}
	return

}

func GetAllCategory() {
	Handlers.GetAllCategoryHandlerConsole()
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		CategoryConsole()
	}
	return
}

func UpdateCategory() {
	Handlers.UpdateCategoryConsole()
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		CategoryConsole()
	} else {
		return
	}
}
