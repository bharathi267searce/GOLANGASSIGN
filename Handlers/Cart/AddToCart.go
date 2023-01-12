package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func AddToCartConsoleHandler(Newcart Datastructures.Cart) {
	var ResponseMessage string
	route := AddToCart(Newcart)

	if route == 441 {
		ResponseMessage = Support.PrepareStatementError
	} else if route == 442 {
		ResponseMessage = Support.OutOfStock
	} else if route == 443 {
		ResponseMessage = Support.UnvailableQuanity
	} else if route == 444 {
		ResponseMessage = Support.NoQuantityCheck
	} else if route == 445 {
		ResponseMessage = Support.InvalidProductId
	} else if route == 446 {
		ResponseMessage = Support.InvalidReferenceId

	} else if route == 201 {
		ResponseMessage = Support.UpdateCartItem
	} else if route == 202 {
		ResponseMessage = Support.InsertedCartItem
	} else if route == 447 {
		ResponseMessage = Support.QuantityMustBePositive
	} else {
		ResponseMessage = Support.ErrorScaning
	}

	Support.PrintResponse(ResponseMessage)
}
func AddToCart(NewCart Datastructures.Cart) int {
	if NewCart.Quantity <= 0 {
		return 447
	}
	rows := Support.CheckProductId(NewCart.Product_id)
	if !rows.Next() {
		return 445
	}

	rows = Support.CheckReferenceId(NewCart.Reference_id)
	//check if valid product is there or not
	if !rows.Next() {
		return 446
	}
	// check if the amount is present inventory or not
	// update inventory

	CheckItemQuantity := CheckQuantity(NewCart.Product_id, NewCart.Quantity)
	if CheckItemQuantity != 200 {
		return CheckItemQuantity
	}

	//if product not present then insert or else update
	if CheckCartProduct(NewCart.Reference_id, NewCart.Product_id, NewCart.Quantity) != 201 {

		insertStatement, err := Support.DB.Prepare(query.AddCart)
		if err != nil {
			return 441
		}
		_, err = insertStatement.Exec(NewCart.Reference_id, NewCart.Product_id, NewCart.Quantity)
		if err != nil {
			fmt.Println("hello2")
			return 441
		}
		fmt.Print("The product added successfully")
		return 202

	}
	return 441

}
func AddToCartRoute(w http.ResponseWriter, r *http.Request) {

	var Newcart Datastructures.Cart
	reqBody, err := ioutil.ReadAll(r.Body)

	var ResponseCode int
	var ResponseMessage string
	if err != nil {
		ResponseCode = Support.NotFound
		ResponseMessage = Support.InvalidCartFormat
	}

	err = json.Unmarshal(reqBody, &Newcart)
	if err != nil {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorUnMarshaling
	}

	route := AddToCart(Newcart)

	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.PrepareStatementError

	} else if route == 442 {
		ResponseCode = Support.NotFound
		ResponseMessage = Support.OutOfStock
	} else if route == 443 {
		ResponseCode = Support.Exsits
		ResponseMessage = Support.UnvailableQuanity
	} else if route == 444 {
		ResponseCode = Support.Error
		ResponseMessage = Support.NoQuantityCheck
	} else if route == 445 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidProductId
	} else if route == 446 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidReferenceId

	} else if route == 201 {
		ResponseCode = Support.Accepted
		ResponseMessage = Support.UpdateCartItem
	} else if route == 202 {
		ResponseCode = Support.Inserted
		ResponseMessage = Support.InsertedCartItem
	} else {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorScaning
	}

	//check if valid product is there or not
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}
