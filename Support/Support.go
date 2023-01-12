package Support

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	query "github.com/1234bharathi/GOLANGASSIGN/Query"
)

// check product id
func CheckProductInventory(Id string) *sql.Rows {
	rows, err := DB.Query(query.GetInventory, Id)
	if err != nil {
		return nil
	}
	return rows
}
func CheckReferenceId(Id string) *sql.Rows {

	rows, err := DB.Query(query.GetReferenceId, Id)
	if err != nil {
		return nil
	}
	return rows
}

// check product id
func CheckProductId(Id string) *sql.Rows {

	rows, err := DB.Query(query.GetProductInfo, Id)
	if err != nil {

		return nil
	}
	// defer rows.Close()
	return rows
}
func WriteResponse(status int, data any, w http.ResponseWriter) {
	// var w http.ResponseWriter
	// w.WriteHeader(http.StatusCreated)
	response := Datastructures.Response{
		Status:  status,
		Message: data,
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

//Page and items checker

func PageNoandItemChecker(PageNo, NoOfItem int) (int, int) {
	if PageNo <= 0 {
		PageNo = 1
	}
	if NoOfItem <= 0 {
		NoOfItem = 20
	}
	return PageNo, NoOfItem
}

// checkCategory_id
func CheckCategoryId(id string) *sql.Rows {
	fmt.Println(id)
	// var category Datastructures.Category_master
	rows, err := DB.Query(query.GetCategory, id)
	// rows.Next()
	// rows.Scan(&category.Category_id, &category.Category_name)
	// fmt.Println(category)

	if err != nil {
		fmt.Print(err)
		return nil
	}
	// defer rows.Close()
	return rows
}

func PrintResponse(data any) {
	fmt.Println(data)
	return
}
