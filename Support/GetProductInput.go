package Support

import (
	"fmt"
	"strings"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
)

func GetProductInput() Datastructures.Product_master {
	var NewProduct Datastructures.Product_master
	fmt.Println("Please enter the valid product id")
	// var product_id string

	_, err := fmt.Scanf("%s", &NewProduct.Product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the product name")

	_, err = fmt.Scanln(&NewProduct.Name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the specification key")
	var m map[string]any
	var SplitString []string

	var SpecificationString string

	_, err = fmt.Scanln(&SpecificationString)
	if err != nil {
		fmt.Println(err)
	}
	if SpecificationString != "" {
		SplitString = strings.Split(SpecificationString, ",")
		m = make(map[string]any)
		for _, pair := range SplitString {
			obj := strings.Split(pair, ":")
			key := obj[0]
			v := strings.Split(obj[1], ",")
			m[key] = v[0]
			fmt.Print(v[0])
		}
		NewProduct.Specification = m
	} else {
		NewProduct.Specification = m
	}

	fmt.Println("Please enter the SKU")

	_, err = fmt.Scanln(&NewProduct.Sku)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("Please enter the Category id")

	_, err = fmt.Scanf("%s", &NewProduct.Category_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the Price")

	_, err = fmt.Scanln(&NewProduct.Price)
	if err != nil {
		fmt.Println(err)
	}
	return NewProduct
}
