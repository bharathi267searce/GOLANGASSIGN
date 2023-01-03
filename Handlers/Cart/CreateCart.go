package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	fmt.Println(name)

	id := uuid.New()
	fmt.Println("Generated UUID:")

	var newcart Datastructures.Cart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
	}

	err = json.Unmarshal(reqBody, &newcart)
	if err != nil {
		fmt.Fprintf(w, "error unmarshalling")
	}
	insertStatement, err := Support.DB.Prepare("INSERT INTO cart_reference(reference_id,create_date,username) VALUES($1,now(),$2)")
	if err != nil {
		fmt.Println("hello2")
		panic(err)
	}
	_, err = insertStatement.Exec(id, name)
	if err != nil {
		fmt.Println("hello2")
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)

}
