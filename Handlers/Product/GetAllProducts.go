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

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	allproducts := []Datastructures.Product_master{}
	page_no, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("page no invalid")
	}
	endlimit := page_no * 2
	startlimit := endlimit - 2
	fmt.Println(startlimit)
	rows, err := Support.DB.Query("Select product_id,name ,price  from product_master ")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var product Datastructures.Product_master
	// var rawContent string
	productlist := []map[string]any{}

	for rows.Next() {
		err := rows.Scan(&product.Product_id, &product.Name, &product.Price)
		if err != nil {
			log.Fatal(err)
		}

		allproducts = append(allproducts, product)
	}

	for _, v := range allproducts {
		newprod := map[string]any{
			"product_id": v.Product_id,
			"name":       v.Name,
			"price":      v.Price,
		}
		productlist = append(productlist, newprod)
	}
	endlimit = int(math.Min(float64(len(allproducts)), float64(endlimit)))
	startlimit = int(math.Min(float64(len(allproducts)-2), float64(startlimit)))

	if startlimit >= 0 && startlimit < endlimit {
		fmt.Println(productlist[startlimit:endlimit])

		result := fmt.Sprintln(productlist[startlimit:endlimit])

		w.Header().Add("Content-Type", "application/json")
		response := Datastructures.Response{
			Status:  http.StatusAccepted,
			Message: result,
		}
		json.NewEncoder(w).Encode(response)
	}

}
