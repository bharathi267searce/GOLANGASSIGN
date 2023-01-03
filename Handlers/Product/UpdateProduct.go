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

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var newproduct Datastructures.Product_master
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
	}

	err = json.Unmarshal(reqBody, &newproduct)
	if err != nil {
		fmt.Println("err in unmarhsling product")
	}
	fmt.Println(newproduct)
	fmt.Println("huuu")
	if !Support.CheckProduct_id(newproduct.Product_id) {
		response := Datastructures.Response{
			Status:  http.StatusForbidden,
			Message: "The value Product_Id Does not exsits, enter valid id",
		}
		json.NewEncoder(w).Encode(response)

		return
	}
	fmt.Println("hilli")
	rows, err := Support.DB.Query("SELECT * from product_master WHERE product_id = $1", newproduct.Product_id)
	if err != nil {
		fmt.Println("err in selecting product")
	}
	defer rows.Close()
	var exsisting_product Datastructures.Product_master
	var rawContent []byte
	for rows.Next() {
		fmt.Println("worki")
		err := rows.Scan(&exsisting_product.Product_id, &exsisting_product.Name, &exsisting_product.Sku, &exsisting_product.Category_id, &exsisting_product.Price, &rawContent)
		if err != nil {
			fmt.Println("error scaning")
			log.Fatal(err)
		}

		err = json.Unmarshal(rawContent, &exsisting_product.Specification)
		if err != nil {
			fmt.Println("error unmarshaling")
		}
		if newproduct.Name == "" {
			newproduct.Name = exsisting_product.Name
		}
		if newproduct.Price == 0 {
			newproduct.Price = exsisting_product.Price
		}
		if newproduct.Sku == "" {
			newproduct.Sku = exsisting_product.Sku
		}
		if newproduct.Specification == nil {
			newproduct.Specification = exsisting_product.Specification
		}
		if newproduct.Category_id == "" {
			newproduct.Category_id = exsisting_product.Category_id
		}
		if newproduct.Category_id != exsisting_product.Category_id {
			// result := fmt.Sprint("Cannot alter the category id for product!!!\n please update in category table")
			response := Datastructures.Response{
				Status:  http.StatusForbidden,
				Message: "Cannot alter the category id for product!!!\n please update in category table",
			}
			json.NewEncoder(w).Encode(response)

			return
		}

		json_specification, err := json.Marshal(newproduct.Specification)

		fmt.Println(newproduct)
		// db.Query("UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;", newproduct.Name, newproduct.Sku, newproduct.Price, json_specification, newproduct.Product_id)
		Support.DB.Query("UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;", newproduct.Name, newproduct.Sku, newproduct.Price, json_specification, newproduct.Product_id)
		if err != nil {
			fmt.Println("error updating")
			response := Datastructures.Response{
				Status:  http.StatusForbidden,
				Message: "error updating",
			}
			json.NewEncoder(w).Encode(response)

			return

		}
		response := Datastructures.Response{
			Status:  http.StatusAccepted,
			Message: "The value Product_Id Does not exsits, enter valid id",
		}
		json.NewEncoder(w).Encode(response)

		return
	}
}
