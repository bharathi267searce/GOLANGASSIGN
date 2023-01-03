package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {

	var newcart Datastructures.Cart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
	}

	err = json.Unmarshal(reqBody, &newcart)
	if err != nil {
		fmt.Fprintf(w, "error unmarshalling")
	}

	//check if valid product is there or not
	if Support.CheckProduct_id(newcart.Product_id) == false {
		result := fmt.Sprintf("The product does not exsits enter a valid product id")
		json.NewEncoder(w).Encode(result)
		return
	}

	//check if valid product is there or not
	if Support.CheckReference_id(newcart.Reference_id) == false {
		result := fmt.Sprintf("The Reference_ID does not exsits enter a valid id")
		json.NewEncoder(w).Encode(result)
		return
	}

	// check if the amount is present inventory or not
	if CheckQuantity(newcart.Product_id, newcart.Quantity) == false {
		result := fmt.Sprintf("OUT OF STOCK!!!The selected quanitity is not available please select less number of items")
		json.NewEncoder(w).Encode(result)
		return
	}

	// update inventory
	//if product not present then insert or else update
	if CheckCartProduct(newcart.Reference_id, newcart.Product_id, newcart.Quantity) == false {
		insertStatement, err := Support.DB.Prepare("INSERT INTO cart(reference_id,product_id,quantity) VALUES($1,$2,$3)")
		if err != nil {
			fmt.Println("hel")
			panic(err)
		}
		_, err = insertStatement.Exec(newcart.Reference_id, newcart.Product_id, newcart.Quantity)
		if err != nil {
			fmt.Println("hello2")
			panic(err)
		}
		result := fmt.Sprint("The product added successfully")
		json.NewEncoder(w).Encode(result)
		w.WriteHeader(http.StatusCreated)

	}

}
