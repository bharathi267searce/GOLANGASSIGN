package Handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
// 	"github.com/1234bharathi/GOLANGASSIGN/Support"
// )

// func AddItem(item Datastructures.Inventory) bool {
// 	var w http.ResponseWriter
// 	//check if valid product is there or not
// 	if !Support.CheckProduct_id(item.Product_id) {
// 		response := Datastructures.Response{
// 			Status:  http.StatusForbidden,
// 			Message: "The product does not exsits enter a valid product id",
// 		}
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(response)
// 		return false
// 	}

// 	// check if the amount is present inventory or not
// 	// fmt.Println("anmm")
// 	fmt.Println(CheckQuantity(item.Product_id, item.Quantity))
// 	if !CheckQuantity(item.Product_id, item.Quantity) {
// 		return false
// 		// fmt.Println("deb")
// 		// response := Datastructures.Response{
// 		// 	Status:  http.StatusForbidden,
// 		// 	Message: "OUT OF STOCK!!!The selected quanitity is not available please select less number of items",
// 		// }
// 		// w.Header().Add("Content-Type", "application/json")
// 		// json.NewEncoder(w).Encode(response)
// 		// return
// 	}
// 	return true
// }
