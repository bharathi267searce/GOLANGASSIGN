package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func GetProductHandlerConsole() {
	var product Datastructures.ProductInfo
	var getproduct_id string
	fmt.Scanln(&getproduct_id)

	route, product := GetProduct(getproduct_id)

	var ResponseMessage any
	if route == 441 {
		ResponseMessage = Support.ExecStatementError

	} else if route == 442 {
		ResponseMessage = Support.ErrorUnMarshalingSpecification
	} else if route == 404 {
		ResponseMessage = Support.InvalidProductId
	} else {
		ResponseMessage = product
	}
	Support.PrintResponse(ResponseMessage)
}
func GetProductRoute(w http.ResponseWriter, r *http.Request) {

	var product Datastructures.ProductInfo
	getproduct_id := mux.Vars(r)["id"]
	route, product := GetProduct(getproduct_id)

	var ResponseCode int
	var ResponseMessage any
	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError

	} else if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorUnMarshalingSpecification
	} else if route == 404 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidProductId
	} else {
		fmt.Println(route)
		ResponseCode = Support.Accepted
		ResponseMessage = product
	}
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}
func GetProduct(ProductId string) (int, Datastructures.ProductInfo) {
	var product Datastructures.ProductInfo
	fmt.Println(ProductId)

	rows := Support.CheckProductId(ProductId)
	if !rows.Next() {
		return 404, product
	}

	rows, err := Support.DB.Query(query.GetProduct, ProductId)
	if err != nil {
		return 441, product
	}
	defer rows.Close()

	var rawContent string
	for rows.Next() {

		err := rows.Scan(&product.Product_id, &product.Name, &product.Sku, &product.Price, &product.CateoryName, &rawContent)
		if err != nil {
			return 441, product
		}
		err = json.Unmarshal([]byte(rawContent), &product.Specification)
		if err != nil {
			return 442, product
		}
		return 202, product

	}
	return 441, product
}
