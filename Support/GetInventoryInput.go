package Support

import (
	"fmt"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
)

func GetInventoryInput() Datastructures.Inventory {
	var Inventory Datastructures.Inventory
	fmt.Println("Please enter the valid Product id")
	// var product_id string

	_, err := fmt.Scanf("%s", &Inventory.Product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter Quantity")

	_, err = fmt.Scanln(&Inventory.Quantity)
	if err != nil {
		fmt.Println(err)
	}

	return Inventory
}
