package Handlers

import (
	"fmt"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func GetCartConsoleHandler() {

	fmt.Println("Enter the ReferenceID")
	var ReferenceID string
	_, err := fmt.Scanf("%s", &ReferenceID)
	var ResponseMessage any
	if err != nil {
		ResponseMessage = Support.ErrorScaningInput
	}

	route, Cart := GetCart(ReferenceID)

	if route == 441 {
		ResponseMessage = Support.InvalidReferenceId
	} else if route == 442 {
		ResponseMessage = Support.ExecStatementError
	} else if route == 443 {
		ResponseMessage = Support.ErrorGetData
	} else {
		ResponseMessage = Cart
	}
	Support.PrintResponse(ResponseMessage)

}
func GetCart(ReferenceId string) (int, any) {

	rows := Support.CheckReferenceId(ReferenceId)
	if !rows.Next() {
		return 441, Support.InvalidReferenceId
	}
	items := []Datastructures.Product_view{}
	rows, err := Support.DB.Query(query.GetCart, ReferenceId)
	if err != nil {
		return 442, Support.ExecStatementError
	}
	defer rows.Close()
	item := Datastructures.Product_view{}
	var sum float32
	for rows.Next() {
		err := rows.Scan(&item.Product_id, &item.Price, &item.Category_name, &item.Quantity)
		if err != nil {
			return 443, Support.ErrorGetData
		}

		sum += item.Price
		items = append(items, item)

	}
	Cart := map[string]interface{}{
		"Total Price": sum,
		"data":        items,
	}
	return 200, Cart

}
func GetCartRoute(w http.ResponseWriter, r *http.Request) {
	ReferenceId := mux.Vars(r)["id"]
	route, Cart := GetCart(ReferenceId)
	var ResponseCode int
	var ResponseMessage any

	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidReferenceId
	} else if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else if route == 443 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorGetData
	} else {
		ResponseCode = Support.Success
		ResponseMessage = Cart
	}
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)

}
