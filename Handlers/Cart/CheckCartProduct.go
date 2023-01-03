package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func CheckCartProduct(Reference_id string, Product_id string, Quantity int) bool {
	rows, err := Support.DB.Query("SELECT quantity from cart WHERE product_id = $1 and reference_id= $2", Product_id, Reference_id)
	if err != nil {
		fmt.Println("errr", err)
	}
	fmt.Println(Reference_id, Product_id, Quantity)
	var w http.ResponseWriter
	for rows.Next() {
		var cartitem int
		rows.Scan(&cartitem)
		cartcount := Quantity + cartitem
		Support.DB.Exec("UPDATE cart SET quantity=$1 WHERE product_id =$2 and reference_id=$3", cartcount, Product_id, Reference_id)
		result := fmt.Sprintln("selected quantity of product added to your cart")
		json.NewEncoder(w).Encode(result)

		return true
	}

	defer rows.Close()
	return false

}
