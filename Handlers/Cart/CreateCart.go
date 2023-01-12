package Handlers

import (
	"fmt"
	"net/http"

	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateCartConsoleHandler() {
	var Name string
	var ResponseMessage string
	fmt.Println("Please enter your Name")
	// var product_id string

	_, err := fmt.Scanf("%s", &Name)
	fmt.Println(Name)
	if err != nil {
		ResponseMessage = Support.ErrorScaningInput
	}
	route := CreateCart(Name)

	if route == 441 {
		ResponseMessage = Support.ErrorGetData

	} else if route == 442 {
		ResponseMessage = Support.ErrorCategoryId
	} else if route == 443 {
		ResponseMessage = Support.InvalidProductId
	}
	Support.PrintResponse(ResponseMessage)

}
func CreateCartRoute(w http.ResponseWriter, r *http.Request) {
	Name := mux.Vars(r)["name"]
	fmt.Println(Name)
	route := CreateCart(Name)

	var ResponseCode int
	var ResponseMessage string
	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorGetData

	} else if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorCategoryId
	} else if route == 443 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidProductId
	}
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)

}
func CreateCart(Name string) int {

	ReferenceId := uuid.New()
	fmt.Println("Generated UUID:")

	insertStatement, err := Support.DB.Prepare(query.CreateCart)
	if err != nil {
		return 441
	}
	_, err = insertStatement.Exec(ReferenceId, Name)
	if err != nil {
		return 442
	}
	return 200

}
