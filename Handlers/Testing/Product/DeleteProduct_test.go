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

func TestDeleteProduct(t *testing.T) {
	// Send the DELETE request
	// http.N
	// http.NewRequest("DELETE", "http://localhost:8081/deleteProduct/ele13", nil)
	req, err := http.NewRequest("DELETE", "http://localhost:8083/deleteProduct/cl19", bytes.NewBuffer(nil))
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
	resp, err = http.Get("http://localhost:8083/getProduct/cl19")
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
	if response.Status != 404 {
		fmt.Println(response.Status)
		fmt.Println("Record was not deleted")
		t.Error("Expected", http.StatusCreated, "Got", response.Status)
		return
	} else {
		fmt.Println("response Status code")
		fmt.Println(response.Status)
		fmt.Println("DELETE request was successful")
	}
}
