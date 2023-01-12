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

func AddInventoryRoute(w http.ResponseWriter, r *http.Request) {
	var Inventory Datastructures.Inventory
	reqBody, err := ioutil.ReadAll(r.Body)
	var ResponseCode int
	var ResponseMessage string
	if err != nil {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidInventoryFormat
	}

	err = json.Unmarshal(reqBody, &Inventory)
	if err != nil {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorUnMarshaling
	}
	fmt.Printf("%+v", Inventory)
	route := AddInventory(Inventory)
	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.InvalidProductId
	} else if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.PrepareStatementError
	} else if route == 443 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else if route == 444 {
		ResponseCode = Support.NotFound
		ResponseMessage = Support.ItemExsistInventory
	} else {
		ResponseCode = Support.Inserted
		ResponseMessage = Support.InventoryInserted
		w.WriteHeader(http.StatusCreated)
	}
	w.Header().Add("Content-Type", "application/json")
	Support.WriteResponse(ResponseCode, ResponseMessage, w)
}
func AddInventory(Inventory Datastructures.Inventory) int {
	rows := Support.CheckProductId(Inventory.Product_id)
	if !rows.Next() {
		return 441
	}
	rows = Support.CheckProductInventory(Inventory.Product_id)

	if !rows.Next() {
		return 444
	}

	insertStatement, err := Support.DB.Prepare(query.AddInventory)
	if err != nil {
		return 442
	}
	_, err = insertStatement.Exec(Inventory.Product_id, Inventory.Quantity)
	if err != nil {
		return 443
	}
	return 200
}

func AddInventoryConsoleHandler(Inventory Datastructures.Inventory) {

	var ResponseMessage string
	fmt.Printf("%+v", Inventory)
	route := AddInventory(Inventory)
	if route == 441 {
		ResponseMessage = Support.InvalidProductId
	} else if route == 442 {
		ResponseMessage = Support.PrepareStatementError
	} else if route == 443 {
		ResponseMessage = Support.ExecStatementError
	} else {
		ResponseMessage = Support.InventoryInserted
	}
	Support.PrintResponse(ResponseMessage)
}
