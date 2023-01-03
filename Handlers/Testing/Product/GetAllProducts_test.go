package Testing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
)

func TestGetAllProducts(t *testing.T) {
	resp, err := http.Get("http://localhost:8083/getProducts/2")
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
		fmt.Println("response.Status")
		fmt.Println(response.Status)
		fmt.Println("response.Message")
		fmt.Println(response.Message)

		fmt.Println("Records Got Sucessfully")

		return
	} else {
		t.Error("Expected", http.StatusCreated, "Got", response.Status)
		return
	}
}
