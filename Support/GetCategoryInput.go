package Support

import (
	"fmt"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
)

func GetCategoryInput() Datastructures.Category_master {
	var Category Datastructures.Category_master
	fmt.Println("Please enter the valid Category id")
	// var product_id string

	_, err := fmt.Scanf("%s", &Category.Category_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the Category name")

	_, err = fmt.Scanln(&Category.Category_name)
	if err != nil {
		fmt.Println(err)
	}

	return Category
}
