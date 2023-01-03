package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]

	res, err := Support.DB.Exec("DELETE FROM cart WHERE reference_id=$1", x)

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				result := fmt.Sprintln("The reference_id Does not exsits, eneter valid id")
				json.NewEncoder(w).Encode(result)
				return
			}
			result := fmt.Sprintln("The cart item is deleted sucessfully")
			json.NewEncoder(w).Encode(result)
		}

	}

	return

}
