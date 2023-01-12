package Support

import (
	"fmt"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
)

func GetCartInput() Datastructures.Cart {
	var Cart Datastructures.Cart
	fmt.Println("Please enter the valid Product ID")
	// var product_id string

	_, err := fmt.Scanf("%s", &Cart.Product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the Reference ID")

	_, err = fmt.Scanln(&Cart.Reference_id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Please enter the Qunantity")

	_, err = fmt.Scanln(&Cart.Quantity)
	if err != nil {
		fmt.Println(err)
	}

	return Cart
}

func GetMultipleCartInput(ReferenceId string) Datastructures.Cart {
	var Cart Datastructures.Cart
	fmt.Println("Please enter the valid Product ID")
	// var product_id string

	_, err := fmt.Scanf("%s", &Cart.Product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the Qunantity")

	_, err = fmt.Scanln(&Cart.Quantity)
	if err != nil {
		fmt.Println(err)
	}
	Cart.Reference_id = ReferenceId
	return Cart
}
