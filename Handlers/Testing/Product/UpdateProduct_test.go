package Testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
)

func TestUpdateProduct(t *testing.T) {
	// Send the DELETE request

	request := Datastructures.Product_master{
		Product_id:  "cl15",
		Name:        "kurti",
		Sku:         "89bnx",
		Category_id: "cl",
		Price:       1000,
		Specification: map[string]any{
			"size":  "xl",
			"color": "green",
			"type":  "loose fit",
		},
	}

	requestBodyJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("PUT", "http://localhost:8083/updateProduct", bytes.NewBuffer(requestBodyJSON))
	// http.DELETE("http://localhost:8081/deleteProduct/ele13")

	if err != nil {

		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// defer req.Body.Close()

	// // Check the status codez
	// if req.StatusCode != 200 || req.StatusCode != 202 {
	// 	fmt.Println("DELETE request was unsuccessful:", resp.StatusCode)
	// 	return
	// }

	// Read and print the response body
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(body))

	// Send a follow-up request to verify that the record was deleted
	fmt.Println("he8")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var response Datastructures.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("error unmarshaling response")
	}
	// Check the status code
	if response.Status == 202 {
		fmt.Println("UPDATE request was successful")

		return
	} else {
		fmt.Println(response.Status)
		fmt.Println("Record was not updated")
		t.Error("Expected", http.StatusAccepted, "Got", response.Status)
	}
}
