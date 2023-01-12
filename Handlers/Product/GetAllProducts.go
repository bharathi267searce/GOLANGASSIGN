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

func GetAllProductsHandlerConsole() {
	var PageNo int
	var ItemsPerPage int
	fmt.Println("Enter the Page Number")
	fmt.Scanln(&PageNo)
	fmt.Println("Enter the Items Per Page")
	fmt.Scanln(&ItemsPerPage)

	PageNo, ItemsPerPage = Support.PageNoandItemChecker(PageNo, ItemsPerPage)

	allproducts := []Datastructures.Product{}
	allproducts, route := GetAllProducts(PageNo, ItemsPerPage)
	var ResponseMessage any
	if route == 441 {

		ResponseMessage = Support.ExecStatementError
	} else if route == 442 {
		ResponseMessage = Support.NotFound
	} else {
		ResponseMessage = allproducts
	}
	Support.PrintResponse(ResponseMessage)

}
func GetAllProducts(PageNo, ItemsPerPage int) ([]Datastructures.Product, int) {

	endlimit := PageNo * ItemsPerPage
	startlimit := endlimit - ItemsPerPage
	allproducts := []Datastructures.Product{}
	fmt.Println(startlimit)
	rows, err := Support.DB.Query(query.GetAllproduct)
	if err != nil {
		return allproducts, 441
	}
	defer rows.Close()

	var product Datastructures.Product
	// var rawContent string
	// productlist := []map[string]any{}

	for rows.Next() {
		err := rows.Scan(&product.Product_id, &product.Name, &product.Price)
		if err != nil {
			return allproducts, 442
		}

		allproducts = append(allproducts, product)
	}

	// for _, v := range allproducts {
	// 	newprod := map[string]any{
	// 		"product_id": v.Product_id,
	// 		"name":       v.Name,
	// 		"price":      v.Price,
	// 	}
	// 	productlist = append(productlist, newprod)
	// }
	endlimit = int(math.Min(float64(len(allproducts)), float64(endlimit)))
	startlimit = int(math.Min(float64(len(allproducts)-2), float64(startlimit)))

	if startlimit >= 0 && startlimit < endlimit {
		// fmt.Println(productlist[startlimit:endlimit])
		return allproducts[startlimit:endlimit], 200
		// result := fmt.Sprintln(productlist[startlimit:endlimit])

		// response := Datastructures.Response{
		// 	Status:  http.StatusAccepted,
		// 	Message: result,
		// }
		// json.NewEncoder(w).Encode(response)
	}
	return allproducts, 442
}
func GetAllProductsRoute(w http.ResponseWriter, r *http.Request) {
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
	allproducts := []Datastructures.Product{}
	fmt.Println(Display.PageNo, Display.ItemsPerPage)
	allproducts, route := GetAllProducts(Display.PageNo, Display.ItemsPerPage)

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
		ResponseMessage = allproducts
	}
	Support.WriteResponse(ResponseCode, ResponseMessage, w)

}
