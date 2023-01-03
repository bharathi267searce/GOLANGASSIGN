package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/1234bharathi/GOLANGASSIGN/Datastructures"
	"github.com/1234bharathi/GOLANGASSIGN/Support"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	var category Datastructures.Category_master
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
	}

	err = json.Unmarshal(reqBody, &category)
	if err != nil {
		fmt.Fprintf(w, "error unmarshalling")
	}
	fmt.Printf("%+v", category)
	insertStatement, err := Support.DB.Prepare("INSERT INTO category_master(category_id,category_name) VALUES($1,$2)")
	if err != nil {
		fmt.Println("hello2")
		panic(err)
	}
	_, err = insertStatement.Exec(category.Category_id, category.Category_name)
	if err != nil {
		fmt.Println("hello2")
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
