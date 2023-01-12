package Handlers

import (
	"fmt"
	"net/http"

	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func DeleteInventoryRoute(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]

	var ResponseCode int
	var ResponseMessage string
	route := DeleteInventory(x)
	if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError

	} else if route == 443 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidProductId
	} else {
		ResponseCode = Support.Accepted
		ResponseMessage = Support.InventoryDeleted
	}
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}
func DeleteInventory(ProductId string) int {

	rows := Support.CheckProductId(ProductId)
	if !rows.Next() {
		return 443
	}
	_, err := Support.DB.Exec(query.DeleteInventory, ProductId)
	if err != nil {
		return 442
	}

	if err != nil {
		return 404
	}
	return 200

}
func DeleteInventoryHandlerConsole() {
	var ProductId string
	var ResponseMessage string
	fmt.Println("Please enter a valid ProductId to delete")

	_, err := fmt.Scanln(&ProductId)
	if err != nil {
		ResponseMessage = Support.ErrorScaning
		Support.PrintResponse(ResponseMessage)
	}

	route := DeleteInventory(ProductId)
	if route == 442 {

		ResponseMessage = Support.ExecStatementError

	} else if route == 443 {
		ResponseMessage = Support.InvalidProductId
	} else {
		ResponseMessage = Support.InventoryDeleted
	}
	Support.PrintResponse(ResponseMessage)
	return
}
