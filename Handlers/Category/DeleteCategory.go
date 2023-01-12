package Handlers

import (
	"fmt"
	"net/http"

	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func DeleteCategoryRoute(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	var ResponseCode int
	var ResponseMessage string
	route := DeleteCategory(x)
	if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError

	} else if route == 443 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidCategoryId
	} else {
		ResponseCode = Support.Accepted
		ResponseMessage = Support.CategoryDeleted
	}
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}

func DeleteCategory(CategoryID string) int {

	rows := Support.CheckCategoryId(CategoryID)
	if !rows.Next() {
		return 443
	}
	_, err := Support.DB.Exec(query.DeleteCategory, CategoryID)
	if err != nil {
		return 442
	}

	if err != nil {
		return 404
	}
	return 200
}
func DeleteCategoryConsoleHandler() {
	fmt.Println("Please enter a valid ProductId to delete")
	var CategoryID string
	var ResponseMessage string
	_, err := fmt.Scanln(&CategoryID)
	if err != nil {
		ResponseMessage = Support.ErrorScaning
		Support.PrintResponse(ResponseMessage)
	}

	route := DeleteCategory(CategoryID)
	if route == 442 {

		ResponseMessage = Support.ExecStatementError

	} else if route == 443 {
		ResponseMessage = Support.InvalidCategoryId
	} else {
		ResponseMessage = Support.CategoryDeleted
	}
	Support.PrintResponse(ResponseMessage)
	return
}
