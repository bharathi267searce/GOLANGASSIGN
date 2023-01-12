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

func UpdateCategoryRoute(w http.ResponseWriter, r *http.Request) {
	var NewCategory Datastructures.Category_master
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Support.WriteResponse(Support.InvalidFormat, Support.InvalidCategoryFormat, w)
	}
	err = json.Unmarshal(reqBody, &NewCategory)
	if err != nil {
		Support.WriteResponse(Support.Error, Support.ErrorUnMarshaling, w)
	}
	route := UpdateCategory(NewCategory)
	var ResponseCode int
	var ResponseMessage string
	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorGetData

	} else if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorCategoryId
	} else if route == 445 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else {
		ResponseCode = Support.Inserted
		ResponseMessage = Support.CategoryUpdated
		// w.WriteHeader(http.StatusCreated)

	}
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}
func UpdateCategory(Category Datastructures.Category_master) int {
	// var category Datastructures.Category_master = Datastructures.Category_master{}
	var exsisting_category Datastructures.Category_master
	// fmt.Print(Category)
	rows := Support.CheckCategoryId(Category.Category_id)
	defer rows.Close()
	if !rows.Next() {
		return 442
	} else {
		err := rows.Scan(&exsisting_category.Category_id, &exsisting_category.Category_name)
		if err != nil {
			return 441
		}
		fmt.Print(Category)

		if Category.Category_name == "" {
			Category.Category_name = exsisting_category.Category_name
		}

		fmt.Println(Category)
		// db.Query("UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;", newproduct.Name, newproduct.Sku, newproduct.Price, json_specification, newproduct.Product_id)
		Support.DB.Query(query.UpdateCategory, Category.Category_name, Category.Category_id)
		if err != nil {
			return 445
		}
		return 200
	}

	return 441

}
func UpdateCategoryConsole() {
	var NewCategory Datastructures.Category_master = Support.GetCategoryInput()
	route := UpdateCategory(NewCategory)
	fmt.Println(NewCategory)
	var ResponseMessage string
	if route == 441 {
		ResponseMessage = Support.ErrorGetData
	} else if route == 442 {
		ResponseMessage = Support.InvalidCategoryId
	} else if route == 445 {
		ResponseMessage = Support.ExecStatementError
	} else {
		ResponseMessage = Support.CategoryUpdated
		// w.WriteHeader(http.StatusCreated)

	}
	Support.PrintResponse(ResponseMessage)

}
