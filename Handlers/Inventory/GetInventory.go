package Handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	getproduct_id := mux.Vars(r)["id"]
	fmt.Println(getproduct_id)
	if !Support.CheckProduct_id(getproduct_id) {
		result := fmt.Sprintln("The product does not exsits enter a valid product id")
		json.NewEncoder(w).Encode(result)
		return
	}

	rows, err := Support.DB.Query(query.GetInventory, getproduct_id)
	if err != nil {
		fmt.Println("err in selecting product")
	}
	defer rows.Close()
	var item Datastructures.Inventory

	for rows.Next() {

		err := rows.Scan(&item.Product_id, &item.Quantity)
		if err != nil {
			fmt.Println("error scaning")
			log.Fatal(err)
		}
		// result := fmt.Sprint("Inventory detais of prodct", product.Product_id, product.Name, product.Sku, product.Price, cateory_name, &product.Specification)
		json.NewEncoder(w).Encode(item)
	}
}
