package Handlers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]

	res, err := Support.DB.Exec("DELETE FROM inventory WHERE product_id=$1", x)

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				result := fmt.Sprint("The value Product_Id Does not exsits, eneter valid id")
				json.NewEncoder(w).Encode(result)
				return
			}
			result := fmt.Sprint("The value is deleted sucessfully")
			json.NewEncoder(w).Encode(result)
		}

	}

	return

}
