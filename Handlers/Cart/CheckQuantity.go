package Handlers

import (
	"fmt"

	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func CheckQuantity(Product_id string, Quantity int) int {
	rows, err := Support.DB.Query(query.CheckQuantity, Product_id)
	if err != nil {
		return 441
	}
	fmt.Println(Product_id, Quantity)
	for rows.Next() {
		var stock int

		rows.Scan(&stock)
		fmt.Println(stock, "stock")
		if stock == 0 {
			fmt.Println("h1")
			return 442

		}
		if stock < Quantity {
			fmt.Println("h2")
			return 443

		} else {
			fmt.Println("h3")
			cartcount := stock - Quantity
			Support.DB.Exec(query.UpdateInventory, cartcount, Product_id)
			return 200
		}
		return 444
	}

	defer rows.Close()
	return 444

}
