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

func AddItemsToCartRoute(w http.ResponseWriter, r *http.Request) {
	Ref_id := mux.Vars(r)["Ref_id"]
	fmt.Println(Ref_id)
	ItemList := []Datastructures.Cart{}
	reqBody, err := ioutil.ReadAll(r.Body)
	var ResponseCode int
	var ResponseMessage string
	if err != nil {
		ResponseCode = Support.NotFound
		ResponseMessage = Support.InvalidCartFormat
	}

	err = json.Unmarshal(reqBody, &ItemList)
	if err != nil {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorUnMarshaling
	}
	for _, v := range ItemList {
		route := AddToCart(v)
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

	}
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}
func AddItemsToCartConsoleHandler() {
	fmt.Print("Enter the Reference ID")
	var ReferenceId string
	fmt.Scanf("%s", &ReferenceId)
	rows := Support.CheckReferenceId(ReferenceId)
	if !rows.Next() {
		fmt.Println(Support.InvalidReferenceId)
		return
	}

	fmt.Println("Enter the number of Items you want to add to cart")
	var NoOfItems int
	fmt.Scanf("%d", &NoOfItems)
	ItemsList := [1000]Datastructures.Cart{}
	for i := 1; i <= NoOfItems; i++ {
		ItemsList[i] = Support.GetMultipleCartInput(ReferenceId)
		fmt.Print(ItemsList[i])
		AddToCartConsoleHandler(ItemsList[i])
	}

}
func AddItemToCart() {

}
