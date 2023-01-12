package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	query "github.com/1234bharathi/GOLANGASSIGN/Query"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func GetAllCategoryHandlerConsole() {
	var PageNo int
	var ItemsPerPage int
	fmt.Println("Enter the Page Number")
	fmt.Scanln(&PageNo)
	fmt.Println("Enter the Items Per Page")
	fmt.Scanln(&ItemsPerPage)
	PageNo, ItemsPerPage = Support.PageNoandItemChecker(PageNo, ItemsPerPage)

	AllCategory := []Datastructures.Category_master{}
	AllCategory, route := GetAllCategory(PageNo, ItemsPerPage)
	var ResponseMessage any
	if route == 441 {

		ResponseMessage = Support.ExecStatementError
	} else if route == 442 {
		ResponseMessage = Support.NotFound
	} else {
		ResponseMessage = AllCategory
	}
	Support.PrintResponse(ResponseMessage)

}
func GetAllCategoryRoute(w http.ResponseWriter, r *http.Request) {

	var Display Datastructures.Display
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Support.WriteResponse(Support.InvalidFormat, Support.InvalidFormatDisplay, w)
	}
	err = json.Unmarshal(reqBody, &Display)
	if err != nil {
		Support.WriteResponse(Support.Error, Support.ErrorUnMarshaling, w)
	}

	Display.PageNo, Display.ItemsPerPage = Support.PageNoandItemChecker(Display.PageNo, Display.ItemsPerPage)
	AllCategory := []Datastructures.Category_master{}

	AllCategory, route := GetAllCategory(Display.PageNo, Display.ItemsPerPage)

	var ResponseCode int
	var ResponseMessage any
	if route == 441 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else if route == 404 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ExecStatementError
	} else if route == 442 {
		ResponseCode = Support.Error
		ResponseMessage = Support.ErrorGetData
	} else {
		ResponseCode = Support.Success
		ResponseMessage = AllCategory
	}
	Support.WriteResponse(ResponseCode, ResponseMessage, w)

}

func GetAllCategory(PageNo, ItemsPerPage int) ([]Datastructures.Category_master, int) {

	endlimit := PageNo * ItemsPerPage
	startlimit := endlimit - ItemsPerPage
	AllCategory := []Datastructures.Category_master{}
	var category Datastructures.Category_master
	rows, err := Support.DB.Query(query.GetAllCategory)
	if err != nil {
		return AllCategory, 441
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&category.Category_id, &category.Category_name)
		if err != nil {
			return AllCategory, 442
		}

		AllCategory = append(AllCategory, category)

	}
	endlimit = int(math.Min(float64(len(AllCategory)), float64(endlimit)))
	if startlimit >= 0 && startlimit < endlimit {
		return AllCategory[startlimit:endlimit], 200
	}
	return AllCategory, 404
}
