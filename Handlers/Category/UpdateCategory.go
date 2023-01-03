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

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category Datastructures.Category_master
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data withid and name only in order to update")
	}
	fmt.Println(reqBody)
	err = json.Unmarshal(reqBody, &category)
	if err != nil {
		fmt.Printf("error marshaling")
	}
	fmt.Println("abababa")
	fmt.Println(category)
	if !Support.CheckCategory_id(category.Category_id) {
		result := fmt.Sprintln("The category does not exsits enter a valid category id")
		json.NewEncoder(w).Encode(result)
		return
	}
	fmt.Println("hilli")
	rows, err := Support.DB.Query("SELECT * from category_master WHERE category_id = $1", category.Category_id)
	if err != nil {
		fmt.Println("err in selecting category")
	}
	defer rows.Close()
	var exsisting_category Datastructures.Category_master
	for rows.Next() {
		fmt.Println("worki")
		err := rows.Scan(&exsisting_category.Category_id, &exsisting_category.Category_name)
		if err != nil {
			fmt.Println("error scaning")
			log.Fatal(err)
		}

		if category.Category_name == "" {
			category.Category_name = exsisting_category.Category_name
		}

		fmt.Println(category)
		// db.Query("UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;", newproduct.Name, newproduct.Sku, newproduct.Price, json_specification, newproduct.Product_id)
		Support.DB.Query("UPDATE category_master SET category_name=$1 WHERE category_id =$1;", category.Category_name, category.Category_id)
		if err != nil {
			fmt.Println("error")
		}
	}
}
