package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	// fmt.Println(x)
	fmt.Println("bhbhbh")

	res, err := Support.DB.Exec("DELETE FROM product_master WHERE product_id=$1", x)

	if err == nil {
		fmt.Println("kjkjk")
		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				// result := fmt.Sprintln("The value Product_Id Does not exsits, eneter valid id")
				// json.NewEncoder(w).Encode(result)
				response := Datastructures.Response{
					Status:  http.StatusForbidden,
					Message: "The value Product_Id Does not exsits, enter valid id",
				}
				json.NewEncoder(w).Encode(response)

				return
			}

		}

	}

	// result := fmt.Sprintln("The value is deleted sucessfully")
	response := Datastructures.Response{
		Status:  http.StatusAccepted,
		Message: "The value is deleted sucessfully",
	}
	json.NewEncoder(w).Encode(response)

}
