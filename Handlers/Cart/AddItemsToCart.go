package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func AddItemsToCart(w http.ResponseWriter, r *http.Request) {
	Ref_id := mux.Vars(r)["Ref_id"]
	fmt.Println(Ref_id)
	itemlist := []Datastructures.Inventory{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Id and Quantity only in order to insert")
	}

	err = json.Unmarshal([]byte(reqBody), &itemlist)
	if err != nil {
		fmt.Println("err un marshalling")
		return
	}

	for i, v := range itemlist {
		fmt.Println(i, v.Product_id, v.Quantity)
		// AddItem(v)
		fmt.Println("debug")
		if Support.CheckProduct_id(v.Product_id) == false {
			result := fmt.Sprintf("The product does not exsits enter a valid product id")
			json.NewEncoder(w).Encode(result)
			return
		}
		if !CheckQuantity(v.Product_id, v.Quantity) {
			response := Datastructures.Response{
				Status:  http.StatusForbidden,
				Message: "Product out of Stock",
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return

		}
		// update inventory
		//if product not present then insert or else update
		if !CheckCartProduct(Ref_id, v.Product_id, v.Quantity) {
			insertStatement, err := Support.DB.Prepare("INSERT INTO cart(reference_id,product_id,quantity) VALUES($1,$2,$3)")
			if err != nil {
				fmt.Println("Error Inserting")
				panic(err)
			}
			_, err = insertStatement.Exec(Ref_id, v.Product_id, v.Quantity)
			if err != nil {
				fmt.Println("Error Inserting")
				panic(err)
			}
			response := Datastructures.Response{
				Status:  http.StatusCreated,
				Message: "Product Added Sucessfully",
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			w.WriteHeader(http.StatusCreated)
			return
		}

	}
}
