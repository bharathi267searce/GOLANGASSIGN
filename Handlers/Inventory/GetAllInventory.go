package Handlers

import (
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func GetAllInventoryRoute(w http.ResponseWriter, r *http.Request) {

	route, AllInventory := GetAllInventory()

	var ResponseCode int
	var ResponseMessage any
	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else {
		ResponseCode = Support.Success
		ResponseMessage = AllInventory
	}
	Support.WriteResponse(ResponseCode, ResponseMessage, w)

}
func GetAllInventoryConsoleHandler() {

	route, AllInventory := GetAllInventory()

	var ResponseMessage any
	if route == 441 {
		ResponseMessage = Support.ExecStatementError
	} else if route == 442 {
		ResponseMessage = Support.ExecStatementError
	} else {
		ResponseMessage = AllInventory
	}
	Support.PrintResponse(ResponseMessage)

}
func GetAllInventory() (int, []Datastructures.Inventory) {
	AllInventory := []Datastructures.Inventory{}
	rows, err := Support.DB.Query(query.GetAllInventory)
	if err != nil {
		return 441, AllInventory
	}
	defer rows.Close()
	var product Datastructures.Inventory
	for rows.Next() {
		err := rows.Scan(&product.Product_id, &product.Quantity)
		if err != nil {
			return 442, AllInventory
		}
		AllInventory = append(AllInventory, product)
	}
	return 200, AllInventory
}
