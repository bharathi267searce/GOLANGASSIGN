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

func TestAddProducts(t *testing.T) {

	requestBody := Datastructures.Product_master{
		Product_id:  "cl19",
		Name:        "Samsung Galaxy watch 5 Pro",
		Sku:         "bh97Jk",
		Category_id: "ele",
		Price:       49999,
		Specification: map[string]any{
			"size":              "14.5mm",
			"weight":            "220 gms",
			"bluetooth version": "5.3",
		},
	}

	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", "http://localhost:8083/addProduct/", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var response Datastructures.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("error unmarshaling response")
	}

	if response.Status == http.StatusCreated {
		fmt.Println(response)
	} else {
		fmt.Println(response)
		t.Error("Expected", http.StatusCreated, "Got", response.Status)
	}
	// if body.Status==http.StatusCreated || resp.Message=="Product addedd sucessfully"{
	// 	fmt.Println("passed")
	// }

}
