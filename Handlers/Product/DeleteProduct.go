package Handlers

import (
	"fmt"
	"net/http"

	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func DeleteProductRoute(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]

	var ResponseCode int
	var ResponseMessage string
	route := DeleteProduct(x)
	if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError

	} else if route == 443 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidProductId
	} else {
		ResponseCode = Support.Accepted
		ResponseMessage = Support.ProductDeleted
	}
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}
func DeleteProduct(ProductId string) int {

	rows := Support.CheckProductId(ProductId)
	if !rows.Next() {
		return 443
	}
	_, err := Support.DB.Exec(query.DeleteProduct, ProductId)
	if err != nil {
		return 442
	}

	if err != nil {
		return 404
	}
	return 200

}
func DeleteProductHandlerConsole() {

	fmt.Println("Please enter a valid ProductId to delete")
	var ProductId string
	var ResponseMessage string
	_, err := fmt.Scanln(&ProductId)
	if err != nil {
		ResponseMessage = Support.ErrorScaning
		Support.PrintResponse(ResponseMessage)
	}

	route := DeleteProduct(ProductId)
	if route == 442 {

		ResponseMessage = Support.ExecStatementError

	} else if route == 443 {
		ResponseMessage = Support.InvalidProductId
	} else {
		ResponseMessage = Support.ProductDeleted
	}
	Support.PrintResponse(ResponseMessage)
	return
}
