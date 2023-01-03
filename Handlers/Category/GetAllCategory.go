package Handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
)

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	AllCategory := []Datastructures.Category_master{}
	page_no, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("page no invalid")
	}
	endlimit := page_no * 20
	startlimit := endlimit - 20
	fmt.Println(startlimit)
	rows, err := Support.DB.Query("Select * from category_master ")
	if err != nil {
		log.Output(1, "error")
	}
	defer rows.Close()
	var category Datastructures.Category_master
	// var rawContent string
	catgorylist := []map[string]any{}

	for rows.Next() {
		err := rows.Scan(&category.Category_id, &category.Category_name)
		if err != nil {
			log.Fatal(err)
		}
		// err = json.Unmarshal([]byte(rawContent), &product.Specification)
		// if err != nil {
		// 	fmt.Println("error marshaling")
		// }

		AllCategory = append(AllCategory, category)

		// result := fmt.Sprintln(product.Price, product.Name)
		// json.NewEncoder(w).Encode(result)
	}
	endlimit = int(math.Min(float64(len(AllCategory)), float64(endlimit)))
	if startlimit >= 0 && startlimit < endlimit {
		fmt.Println(AllCategory[startlimit:endlimit])
	}
	for _, v := range AllCategory {
		newcategory := map[string]any{
			" id":  v.Category_id,
			"name": v.Category_name,
		}
		catgorylist = append(catgorylist, newcategory)
	}
	json.NewEncoder(w).Encode(catgorylist)

}
