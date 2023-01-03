package Handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func GetAllInventory(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	allproducts := []Datastructures.Inventory{}
	page_no, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("page ni invalid")
	}
	endlimit := page_no * 20
	startlimit := endlimit - 20
	fmt.Println(startlimit)
	rows, err := Support.DB.Query("Select *  from inventory ")
	defer rows.Close()
	var product Datastructures.Inventory
	// var rawContent string
	productlist := []map[string]any{}

	for rows.Next() {
		err := rows.Scan(&product.Product_id, &product.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		// err = json.Unmarshal([]byte(rawContent), &product.Specification)
		// if err != nil {
		// 	fmt.Println("error marshaling")
		// }

		allproducts = append(allproducts, product)

		// result := fmt.Sprintln(product.Price, product.Name)
		// json.NewEncoder(w).Encode(result)
	}
	endlimit = int(math.Min(float64(len(allproducts)), float64(endlimit)))
	if startlimit >= 0 && startlimit < endlimit {
		fmt.Println(allproducts[startlimit:endlimit])
	}
	for _, v := range allproducts {
		newprod := map[string]any{
			"Product_id": v.Product_id,
			"Quantity":   v.Quantity,
		}
		productlist = append(productlist, newprod)
	}
	json.NewEncoder(w).Encode(allproducts)

}
