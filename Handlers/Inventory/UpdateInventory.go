package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	var item Datastructures.Inventory
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the product_id and name only in order to insert")
	}

	err = json.Unmarshal(reqBody, &item)
	fmt.Println(item)
	if Support.CheckProduct_id(item.Product_id) == false {
		result := fmt.Sprintf("The product does not exsits enter a valid product id")
		json.NewEncoder(w).Encode(result)
		return
	}
	fmt.Println("hilli")
	rows, err := Support.DB.Query("SELECT * from inventory WHERE product_id = $1", item.Product_id)
	if err != nil {
		fmt.Println("err in selecting product")
	}
	defer rows.Close()
	var exsisting_product Datastructures.Inventory
	for rows.Next() {
		fmt.Println("worki")
		err := rows.Scan(&exsisting_product.Product_id, &exsisting_product.Quantity)
		if err != nil {
			fmt.Println("error scaning")
			log.Fatal(err)
		}
		if item.Quantity <= 0 {
			item.Quantity = 0
		}

		fmt.Println(item)
		// db.Query("UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;", newproduct.Name, newproduct.Sku, newproduct.Price, json_specification, newproduct.Product_id)
		Support.DB.Query("UPDATE inventory SET quantity=$1 WHERE product_id =$2", item.Quantity, item.Product_id)
		if err != nil {
			fmt.Println("")
		}
	}
}
