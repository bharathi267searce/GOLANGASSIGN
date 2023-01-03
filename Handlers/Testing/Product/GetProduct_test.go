package Testing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
)

func TestGetProduct(t *testing.T) {
	// Send the DELETE request
	// // http.N
	// // http.NewRequest("DELETE", "http://localhost:8081/deleteProduct/ele13", nil)
	// req, err := http.NewRequest("DELETE", "http://localhost:8082/deleteProduct/cl19", bytes.NewBuffer(nil))
	// // http.DELETE("http://localhost:8081/deleteProduct/ele13")

	// if err != nil {

	// 	fmt.Println(err)
	// 	return
	// }

	// req.Header.Set("Content-Type", "application/json")

	// // Send the request
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()

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

	// Send a request to verify that the record was deleted
	fmt.Println("he8")
	resp, err := http.Get("http://localhost:8083/getProduct/cl15")
	if err != nil {
		fmt.Println("he8p")
		fmt.Println(err)
		return
	}
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
		fmt.Println(response.Status)
		fmt.Println(response.Message)

		fmt.Println("Record Got Sucessfully")

		return
	} else {
		t.Error("Expected", http.StatusCreated, "Got", response.Status)
		return
	}
}
