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

func UpdateInventoryRoute(w http.ResponseWriter, r *http.Request) {
	var NewInventory Datastructures.Inventory
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Support.WriteResponse(Support.InvalidFormat, Support.InvalidInventoryFormat, w)
	}
	err = json.Unmarshal(reqBody, &NewInventory)
	if err != nil {
		Support.WriteResponse(Support.Error, Support.ErrorUnMarshaling, w)
	}
	route := UpdateInventory(NewInventory)
	var ResponseCode int
	var ResponseMessage string
	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorGetData

	} else if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidProductId
	} else if route == 445 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else {
		ResponseCode = Support.Inserted
		ResponseMessage = Support.InventoryUpdated
		// w.WriteHeader(http.StatusCreated)

	}
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}
func UpdateInventory(Product Datastructures.Inventory) int {
	var ExsistingInventory Datastructures.Inventory
	rows := Support.CheckProductId(Product.Product_id)
	defer rows.Close()
	if !rows.Next() {
		return 442
	} else {
		rows, err := Support.DB.Query(query.GetInventory, Product.Product_id)
		if err != nil {
			return 441
		}
		rows.Next()
		err = rows.Scan(&ExsistingInventory.Product_id, &ExsistingInventory.Quantity)
		if err != nil {
			return 441
		}
		fmt.Print(ExsistingInventory)

		if Product.Quantity <= 0 {
			Product.Quantity = ExsistingInventory.Quantity
		}

		fmt.Println(Product)
		// db.Query("UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;", newproduct.Name, newproduct.Sku, newproduct.Price, json_specification, newproduct.Product_id)
		Support.DB.Query(query.UpdateInventory, Product.Quantity, Product.Product_id)
		if err != nil {
			return 445
		}
		return 200
	}

	return 441

}
func UpdateInventoryConsole() {
	var NewInventory Datastructures.Inventory = Support.GetInventoryInput()
	route := UpdateInventory(NewInventory)
	fmt.Println(NewInventory)
	var ResponseMessage string
	if route == 441 {
		ResponseMessage = Support.ErrorGetData
	} else if route == 442 {
		ResponseMessage = Support.InvalidProductId
	} else if route == 445 {
		ResponseMessage = Support.ExecStatementError
	} else {
		ResponseMessage = Support.InventoryUpdated
		// w.WriteHeader(http.StatusCreated)
	}
	Support.PrintResponse(ResponseMessage)
}
