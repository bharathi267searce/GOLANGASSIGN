package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func AddInventory(w http.ResponseWriter, r *http.Request) {
	var inventory Datastructures.Inventory
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
	}

	err = json.Unmarshal(reqBody, &inventory)
	if err != nil {
		fmt.Fprintf(w, "error unmarshalling")
	}
	fmt.Printf("%+v", inventory)
	if Support.CheckProduct_id(inventory.Product_id) == false {
		result := fmt.Sprintln("The product does not exsits enter a valid product id")
		json.NewEncoder(w).Encode(result)
		return
	}
	insertStatement, err := Support.DB.Prepare("INSERT INTO inventory(product_id,quantity) VALUES($1,$2)")
	if err != nil {
		fmt.Println("hello2")
		panic(err)
	}
	_, err = insertStatement.Exec(inventory.Product_id, inventory.Quantity)
	if err != nil {
		fmt.Println("hello2")
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(inventory)
}
