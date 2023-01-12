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

func UpdateProductRoute(w http.ResponseWriter, r *http.Request) {
	var newproduct Datastructures.Product_master
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Support.WriteResponse(Support.InvalidFormat, Support.InvalidProductFormat, w)
	}
	err = json.Unmarshal(reqBody, &newproduct)
	if err != nil {
		Support.WriteResponse(Support.Error, Support.ErrorUnMarshaling, w)
	}
	route := UpdateProduct(newproduct)
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
	} else if route == 444 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidCategoryId
	} else if route == 445 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else if route == 446 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorUnMarshalingSpecification
	} else {
		ResponseCode = Support.Inserted
		ResponseMessage = Support.ProductUpdated
		w.WriteHeader(http.StatusCreated)

	}
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)

	fmt.Println(newproduct)

}
func UpdateProductConsole() {
	var NewProduct Datastructures.Product_master = Support.GetProductInput()
	route := UpdateProduct(NewProduct)
	fmt.Println(NewProduct)
	var ResponseMessage string
	if route == 441 {
		ResponseMessage = Support.ErrorGetData
	} else if route == 442 {
		ResponseMessage = Support.ErrorCategoryId
	} else if route == 443 {
		ResponseMessage = Support.InvalidProductId
	} else if route == 444 {
		ResponseMessage = Support.InvalidCategoryId
	} else if route == 445 {
		ResponseMessage = Support.ExecStatementError
	} else if route == 446 {
		ResponseMessage = Support.ErrorUnMarshalingSpecification
	} else {
		ResponseMessage = Support.ProductUpdated
	}
	Support.PrintResponse(ResponseMessage)

}
func UpdateProduct(NewProduct Datastructures.Product_master) int {
	rows := Support.CheckProductId(NewProduct.Product_id)
	if !rows.Next() {
		return 443
	}
	//check category id
	rows = Support.CheckCategoryId(NewProduct.Category_id)
	if !rows.Next() {
		return 444
	}
	defer rows.Close()
	fmt.Println("hilli")
	rows, err := Support.DB.Query(query.GetProductInfo, NewProduct.Product_id)
	if err != nil {
		return 445
	}
	defer rows.Close()
	var exsisting_product Datastructures.Product_master
	var rawContent []byte
	for rows.Next() {
		err := rows.Scan(&exsisting_product.Product_id, &exsisting_product.Name, &exsisting_product.Sku, &exsisting_product.Category_id, &exsisting_product.Price, &rawContent)
		if err != nil {
			return 441
		}

		err = json.Unmarshal(rawContent, &exsisting_product.Specification)
		if err != nil {
			return 446
		}
		if NewProduct.Name == "" {
			NewProduct.Name = exsisting_product.Name
		}
		if NewProduct.Price == 0 {
			NewProduct.Price = exsisting_product.Price
		}
		if NewProduct.Sku == "" {
			NewProduct.Sku = exsisting_product.Sku
		}
		if NewProduct.Specification == nil {
			NewProduct.Specification = exsisting_product.Specification
		}
		if NewProduct.Category_id == "" {
			NewProduct.Category_id = exsisting_product.Category_id
		}
		if NewProduct.Category_id != exsisting_product.Category_id {
			return 444
		}

		json_specification, err := json.Marshal(NewProduct.Specification)

		fmt.Println(NewProduct)
		Support.DB.Query(query.UpdateProduct, NewProduct.Name, NewProduct.Sku, NewProduct.Price, json_specification, NewProduct.Product_id)
		if err != nil {
			return 445
		}
		return 200
	}
	return 441
}
