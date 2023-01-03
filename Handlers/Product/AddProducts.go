package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func AddProducts(w http.ResponseWriter, r *http.Request) {
	var newproduct Datastructures.Product_master
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
	}

	err = json.Unmarshal(reqBody, &newproduct)
	if err != nil {
		fmt.Fprintf(w, "error unmarshalling")
	}
	fmt.Printf("%+v", newproduct)

	insertStatement, err := Support.DB.Prepare("INSERT INTO product_master(product_id,name,sku,category_id,price,specification) VALUES($1,$2,$3,$4,$5,$6)")
	if err != nil {
		fmt.Println("hello2")
		panic(err)
	}
	json_specification, err := json.Marshal(newproduct.Specification)
	if err != nil {
		fmt.Println("Error Marshaling")
		panic(err)
	}
	_, err = insertStatement.Exec(newproduct.Product_id, newproduct.Name, newproduct.Sku, newproduct.Category_id, newproduct.Price, json_specification)
	if err != nil {
		fmt.Println("Error mkmkm")
		response := Datastructures.Response{
			Status:  http.StatusForbidden,
			Message: "Product canot be added already exsist",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
		panic(err)

	} else {
		w.WriteHeader(http.StatusCreated)
		response := Datastructures.Response{
			Status:  http.StatusCreated,
			Message: "Product addedd sucessfully",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}
