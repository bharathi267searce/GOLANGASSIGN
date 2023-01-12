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

func AddCategoryConsoleHandle(Category Datastructures.Category_master) {

	route := AddCategory(Category)
	var ResponseMessage string
	if route == 441 {
		ResponseMessage = Support.PrepareStatementError

	} else if route == 442 {
		ResponseMessage = Support.ErrorUnMarshalingSpecification
	} else if route == 444 {
		ResponseMessage = Support.CategoryExsist
	} else if route == 445 {
		ResponseMessage = Support.ExecStatementError
	} else {
		ResponseMessage = Support.CategoryInserted

	}
	Support.PrintResponse(ResponseMessage)
}
func AddCategory(Category Datastructures.Category_master) int {
	insertStatement, err := Support.DB.Prepare(query.AddCategory)
	if err != nil {
		return 441
	}
	rows := Support.CheckCategoryId(Category.Category_id)
	if rows.Next() {
		return 444
	}
	rows.Close()

	_, err = insertStatement.Exec(Category.Category_id, Category.Category_name)
	if err != nil {
		return 445
	}
	return 200

}

func AddCategoryRoute(w http.ResponseWriter, r *http.Request) {
	var Category Datastructures.Category_master
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Support.WriteResponse(Support.InvalidFormat, Support.InvalidCategoryFormat, w)
	}
	err = json.Unmarshal(reqBody, &Category)
	if err != nil {
		Support.WriteResponse(Support.Error, Support.ErrorUnMarshaling, w)
	}
	fmt.Printf("%+v", Category)

	route := AddCategory(Category)
	var ResponseCode int
	var ResponseMessage string
	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.PrepareStatementError

	} else if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorUnMarshalingSpecification
	} else if route == 443 {
		ResponseCode = Support.Exsits
		ResponseMessage = Support.CategoryExsist
	} else if route == 445 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else {
		ResponseCode = Support.Inserted
		ResponseMessage = Support.CategoryInserted
		w.WriteHeader(http.StatusCreated)

	}
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}
