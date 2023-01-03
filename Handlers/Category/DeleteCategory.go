package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]

	res, err := Support.DB.Exec("DELETE FROM category_master WHERE category_id=$1", x)

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				result := fmt.Sprintln("The value category_id Does not exsits, eneter valid id")
				json.NewEncoder(w).Encode(result)
				return
			}

		}

	}
	result := fmt.Sprintln("The value is deleted sucessfully")

	json.NewEncoder(w).Encode(result)
	// return

}
