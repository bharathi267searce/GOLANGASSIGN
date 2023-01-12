package ConsoleInterface

import (
	"fmt"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"

	Handlers "github.com/1234bharathi/GOLANGASSIGN/Handlers/Product"
)

func ProductConsole() {
	fmt.Println("Hi, you are here to perform CRUD operations on 'product' table")
	fmt.Println("Please choose the task to perform")
	fmt.Printf("1.Insert\n2.Read with product id\n3.Delete\n4.Get All Products by page number  \n5.Update\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		InsertProduct()
	} else if choice == 2 {
		GetProduct()
	} else if choice == 3 {
		DeleteProduct()
	} else if choice == 4 {
		GetAllProduct()
	} else if choice == 5 {
		UpdateProduct()
	}

}

func DeleteProduct() {
	Handlers.DeleteProductHandlerConsole()
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		ProductConsole()
	}
	return

}
func InsertProduct() {
	var NewProduct Datastructures.Product_master = Support.GetProductInput()

	Handlers.AddProductConsoleHandle(NewProduct)
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		ProductConsole()
	} else {
		return
	}
}

func GetProduct() {
	Handlers.GetProductHandlerConsole()

	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		ProductConsole()
	}
	return

}
func GetAllProduct() {
	Handlers.GetAllProductsHandlerConsole()
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		ProductConsole()
	}
	return
}

func UpdateProduct() {
	Handlers.UpdateProductConsole()
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		ProductConsole()
	} else {
		return
	}
}
