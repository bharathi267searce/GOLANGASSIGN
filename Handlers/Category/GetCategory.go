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

func GetCategory(w http.ResponseWriter, r *http.Request) {
	getCategory_id := mux.Vars(r)["id"]
	fmt.Println(getCategory_id)
	if !Support.CheckCategory_id(getCategory_id) {
		result := fmt.Sprintln("The product does not exsits enter a valid product id")
		json.NewEncoder(w).Encode(result)
		return
	}
	rows, err := Support.DB.Query("SELECT * from category_master  WHERE category_id = $1", getCategory_id)
	if err != nil {
		fmt.Println("err in selecting category")
	}
	defer rows.Close()
	var category Datastructures.Category_master

	for rows.Next() {

		err := rows.Scan(&category.Category_id, &category.Category_name)
		if err != nil {
			fmt.Println("error scaning")
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(category)
	}
}
