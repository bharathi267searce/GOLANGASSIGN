package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func CheckCartProduct(Reference_id string, Product_id string, Quantity int) int {
	rows, err := Support.DB.Query(query.CheckCartItem, Product_id, Reference_id)
	if err != nil {
		return 441
	}
	fmt.Println(Reference_id, Product_id, Quantity)
	var w http.ResponseWriter
	for rows.Next() {
		var cartitem int
		rows.Scan(&cartitem)
		cartcount := Quantity + cartitem
		Support.DB.Exec(query.UpdateCartItem, cartcount, Product_id, Reference_id)
		result := fmt.Sprintln("selected quantity of product added to your cart")
		json.NewEncoder(w).Encode(result)

		return 201
	}

	defer rows.Close()
	return 441

}
