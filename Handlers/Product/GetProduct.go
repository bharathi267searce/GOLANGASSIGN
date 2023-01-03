package Handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	getproduct_id := mux.Vars(r)["id"]
	fmt.Println(getproduct_id)
	if Support.CheckProduct_id(getproduct_id) == false {
		response := Datastructures.Response{
			Status:  http.StatusNotFound,
			Message: "Id not present",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	rows, err := Support.DB.Query("SELECT product_master.product_id,product_master.name,product_master.sku,product_master.price,category_master.category_name,product_master.specification  from product_master JOIN category_master ON product_master.category_id = category_master.category_id WHERE product_id = $1", getproduct_id)
	if err != nil {
		fmt.Println("err in selecting product")
	}
	defer rows.Close()
	var product Datastructures.Product_master
	var rawContent string
	var cateory_name string
	for rows.Next() {

		err := rows.Scan(&product.Product_id, &product.Name, &product.Sku, &product.Price, &cateory_name, &rawContent)
		if err != nil {
			fmt.Println("error scaning")
			log.Fatal(err)
		}
		err = json.Unmarshal([]byte(rawContent), &product.Specification)
		if err != nil {
			fmt.Println("error unmarshaling")
		}

		result := fmt.Sprint("details of product", product.Product_id, product.Name, product.Sku, product.Price, cateory_name, &product.Specification)
		response := Datastructures.Response{
			Status:  http.StatusAccepted,
			Message: result,
		}
		json.NewEncoder(w).Encode(response)

	}
}
