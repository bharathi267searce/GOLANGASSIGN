package Handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	if Support.CheckReference_id(x) == false {
		response := Datastructures.Response{
			Status:  http.StatusForbidden,
			Message: "The Reference_id does not exsits enter a valid  id",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	items := []Datastructures.Product_view{}
	rows, err := Support.DB.Query("select d.product_id,p.price,c.category_name,d.quantity from product_master p  inner join  category_master c on c.category_id=p.category_id inner join cart d on d.product_id = p.product_id where d.reference_id =$1 ", x)
	if err != nil {
		fmt.Println("page no invalid")
	}
	defer rows.Close()
	item := Datastructures.Product_view{}
	// itemlist := []map[string]any{}

	for rows.Next() {
		err := rows.Scan(&item.Product_id, &item.Price, &item.Category_name, &item.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		// err = json.Unmarshal([]byte(rawContent), &product.Specification)
		// if err != nil {
		// 	fmt.Println("error marshaling")
		// }

		items = append(items, item)

	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)

}
