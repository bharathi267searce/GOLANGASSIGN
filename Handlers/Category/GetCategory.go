package Handlers

import (
	"fmt"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func GetCategoryConsoleHandler() {

	var getCategoryID string
	fmt.Println("Enter the Category ID")
	fmt.Scanln(&getCategoryID)
	var Category Datastructures.Category_master
	route, Category := GetCategory(getCategoryID)

	var ResponseMessage any
	if route == 441 {
		ResponseMessage = Support.ErrorGetData

	} else if route == 404 {
		ResponseMessage = Support.InvalidCategoryId
	} else {
		ResponseMessage = Category
	}
	Support.PrintResponse(ResponseMessage)

}
func GetCategoryRoute(w http.ResponseWriter, r *http.Request) {
	GetCategoryId := mux.Vars(r)["id"]
	var Category Datastructures.Category_master
	route, Category := GetCategory(GetCategoryId)
	var ResponseCode int
	var ResponseMessage any
	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorGetData

	} else if route == 404 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidCategoryId
	} else {
		fmt.Println(route)
		ResponseCode = Support.Accepted
		ResponseMessage = Category
	}
	Support.WriteResponse(ResponseCode, ResponseMessage, w)

}
func GetCategory(GetCategoryId string) (int, Datastructures.Category_master) {
	var category Datastructures.Category_master
	fmt.Println(GetCategoryId)

	rows := Support.CheckCategoryId(GetCategoryId)

	if !rows.Next() {
		return 404, category
	} else {
		err := rows.Scan(&category.Category_name, &category.Category_id)
		if err != nil {
			return 441, category
		}
		fmt.Println(category.Category_name)
		return 200, category

	}

	return 200, category

}
