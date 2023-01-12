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

func AddProductRoute(w http.ResponseWriter, r *http.Request) {
	var newproduct Datastructures.Product_master
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Support.WriteResponse(Support.InvalidFormat, Support.InvalidProductFormat, w)
	}
	err = json.Unmarshal(reqBody, &newproduct)
	if err != nil {
		Support.WriteResponse(Support.Error, Support.ErrorUnMarshaling, w)
	}
	fmt.Printf("%+v", newproduct)

	route := AddProducts(newproduct)
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
		ResponseMessage = Support.ProductExsist
	} else if route == 444 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidCategoryId
	} else if route == 445 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else {
		ResponseCode = Support.Inserted
		ResponseMessage = Support.ProductInserted
		w.WriteHeader(http.StatusCreated)

	}
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}

func AddProducts(newproduct Datastructures.Product_master) int {

	insertStatement, err := Support.DB.Prepare(query.AddProduct)

	if err != nil {
		fmt.Println(err)
		return 441
	}
	json_specification, err := json.Marshal(newproduct.Specification)
	if err != nil {
		return 442
	}
	//check product id
	rows := Support.CheckProductId(newproduct.Product_id)
	if rows.Next() {
		return 443
	}
	//check category id
	rows = Support.CheckCategoryId(newproduct.Category_id)
	if !rows.Next() {
		fmt.Print(rows.Scan(newproduct.Category_id, newproduct.Name))
		return 444
	}
	rows.Close()
	_, err = insertStatement.Exec(newproduct.Product_id, newproduct.Name, newproduct.Sku, newproduct.Category_id, newproduct.Price, json_specification)
	if err != nil {
		return 445
	} else {
		return 200
	}
}
func AddProductConsoleHandle(newproduct Datastructures.Product_master) {
	route := AddProducts(newproduct)
	var ResponseMessage string
	if route == 441 {
		ResponseMessage = Support.PrepareStatementError
	} else if route == 442 {
		ResponseMessage = Support.ErrorMarshalingSpecification
	} else if route == 443 {
		ResponseMessage = Support.ProductExsist
	} else if route == 444 {
		ResponseMessage = Support.InvalidCategoryId
	} else if route == 445 {
		ResponseMessage = Support.PrepareStatementError
	} else {
		ResponseMessage = Support.ProductInserted
	}
	Support.PrintResponse(ResponseMessage)
}
