package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func DeleteCartRoute(w http.ResponseWriter, r *http.Request) {

	var ReferenceId string
	var ProductId string
	var ResponseCode int
	var ResponseMessage string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseCode = Support.NotFound
		ResponseMessage = Support.InvalidCartFormat
	}

	err = json.Unmarshal(reqBody, &ReferenceId)

	if err != nil {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorUnMarshaling
	}

	route := DeleteCart(ReferenceId, ProductId)

	if route == 441 {
		ResponseCode = Support.NotFound
		ResponseMessage = Support.ExecStatementError
	} else if route == 442 {
		ResponseCode = Support.NotFound
		ResponseMessage = Support.InvalidReferenceId
	} else if route == 443 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidProductId
	} else {
		ResponseCode = Support.Accepted
		ResponseMessage = Support.CartDeleted
	}
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}
func DeleteCart(ReferenceId, ProductId string) int {
	rows := Support.CheckReferenceId(ReferenceId)
	if !rows.Next() {
		return 442
	}
	rows.Close()
	rows = Support.CheckProductId(ProductId)
	if !rows.Next() {
		return 443
	}
	_, err := Support.DB.Exec(query.DeleteCart, ReferenceId, ProductId)
	if err != nil {
		return 441
	}

	rows.Close()

	return 200

}
func DeleteCartConsoleHandler() {
	fmt.Println("Enter the Reference ID")
	var ReferenceId string
	_, err := fmt.Scanf("%s", &ReferenceId)
	if err != nil {
		fmt.Println(err)
		fmt.Println(Support.ErrorScaningInput)
	}
	fmt.Println("Enter the product ID")
	var ProductId string
	_, err = fmt.Scanf("%s", &ProductId)
	if err != nil {
		fmt.Println(err)
		fmt.Println(Support.ErrorScaningInput)
	}

	route := DeleteCart(ReferenceId, ProductId)
	var ResponseMessage string
	if route == 441 {
		ResponseMessage = Support.ExecStatementError
	} else if route == 442 {
		ResponseMessage = Support.InvalidReferenceId
	} else if route == 443 {
		ResponseMessage = Support.InvalidProductId
	} else {
		ResponseMessage = Support.CartDeleted
	}
	Support.PrintResponse(ResponseMessage)
}
