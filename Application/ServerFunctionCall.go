package application

import (
	"fmt"
	"log"
	"net/http"

	Cart "github.com/1234bharathi/GOLANGASSIGN/Handlers/Cart"
	Category "github.com/1234bharathi/GOLANGASSIGN/Handlers/Category"
	Inventory "github.com/1234bharathi/GOLANGASSIGN/Handlers/Inventory"
	Product "github.com/1234bharathi/GOLANGASSIGN/Handlers/Product"
	"github.com/gorilla/mux"
)

func ServerFunctionCall() {

	router := mux.NewRouter()

	fmt.Println("Hi navigate to postman")
	router.HandleFunc("/addProduct/", Product.AddProductRoute).Methods("POST")
	router.HandleFunc("/getProduct/{id}", Product.GetProductRoute).Methods("GET")
	router.HandleFunc("/getProducts", Product.GetAllProductsRoute).Methods("GET")
	router.HandleFunc("/deleteProduct/{id}", Product.DeleteProductRoute).Methods("DELETE")
	router.HandleFunc("/updateProduct", Product.UpdateProductRoute).Methods("PUT")
	// category handlers call
	router.HandleFunc("/addCategory", Category.AddCategoryRoute).Methods("POST")
	router.HandleFunc("/getCategory/{id}", Category.GetCategoryRoute).Methods("GET")
	router.HandleFunc("/getAllCategory/{id}", Category.GetAllCategoryRoute).Methods("GET")
	router.HandleFunc("/deleteCategory/{id}", Category.DeleteCategoryRoute).Methods("DELETE")
	router.HandleFunc("/updateCategory", Category.UpdateCategoryRoute).Methods("PUT")
	// inventory handlers call
	router.HandleFunc("/addInventory", Inventory.AddInventoryRoute).Methods("POST")
	router.HandleFunc("/getInventory/{id}", Inventory.GetInventory).Methods("GET")
	router.HandleFunc("/getAllInventory", Inventory.GetAllInventoryRoute).Methods("GET")
	router.HandleFunc("/deleteInventory/{id}", Inventory.DeleteInventoryRoute).Methods("DELETE")
	router.HandleFunc("/updateInventory", Inventory.UpdateInventoryRoute).Methods("PUT")

	//cart handlers

	router.HandleFunc("/addToCart", Cart.AddToCartRoute).Methods("POST")
	router.HandleFunc("/addItemsToCart/{Ref_id}", Cart.AddItemsToCartRoute).Methods("POST")
	router.HandleFunc("/createCart/{name}", Cart.CreateCartRoute).Methods("POST")

	router.HandleFunc("/getCart/{id}", Cart.GetCart).Methods("GET")
	router.HandleFunc("/deleteCart/{id}", Cart.DeleteCartRoute).Methods("DELETE")
	// router.HandleFunc("/updateCart", updateCart).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8084", router))
}
