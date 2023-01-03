package main

import (
	"log"
	"net/http"

	// "github.com/1234bharathi/GOLANGASSIGN/Handlers"
	Cart "github.com/1234bharathi/GOLANGASSIGN/Handlers/Cart"
	Category "github.com/1234bharathi/GOLANGASSIGN/Handlers/Category"
	Inventory "github.com/1234bharathi/GOLANGASSIGN/Handlers/Inventory"
	Product "github.com/1234bharathi/GOLANGASSIGN/Handlers/Product"

	"github.com/1234bharathi/GOLANGASSIGN/Support"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// var db *sql.DB

// func DbConnect() {
// 	var err error
// 	connStr := "postgres://postgres:anbu@localhost/postgres?sslmode=disable"
// 	db, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err = db.Ping(); err != nil {
// 		fmt.Println("hlo")
// 		panic(err)
// 	}
// 	fmt.Printf("\nSuccessfully connected to database!\n")
// }

// func HomeLink(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "home page")
// }

// // capital letter for json
// type Product_master struct {
// 	Product_id    string         `json:"product_id"`
// 	Name          string         `json:"name"`
// 	Sku           string         `json:"sku"`
// 	Category_id   string         `json:"category_id"`
// 	Price         float64        `json:"price"`
// 	Specification map[string]any `json:"specification"`
// }
// type Inventory struct {
// 	Product_id string `json:"product_id"`
// 	Quantity   int    `json:"quantity"`
// }
// type Category_master struct {
// 	Category_id   string `json:"category_id"`
// 	Category_name string `json:"category_name"`
// }

// type Cart_reference struct {
// 	Reference_id string    `json:"reference_id"`
// 	Name         string    `json:"name"`
// 	Date         time.Time `json:"date"`
// }

// type Cart struct {
// 	Reference_id string `json:"reference_id"`
// 	Product_id   string `json:"product_id"`
// 	Quantity     int    `json:"quantity"`
// }

// func AddProducts(w http.ResponseWriter, r *http.Request) {
// 	var newproduct Product_master
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
// 	}

// 	err = json.Unmarshal(reqBody, &newproduct)
// 	if err != nil {
// 		fmt.Fprintf(w, "error unmarshalling")
// 	}
// 	fmt.Printf("%+v", newproduct)
// 	insertStatement, err := db.Prepare("INSERT INTO product_master(product_id,name,sku,category_id,price,specification) VALUES($1,$2,$3,$4,$5,$6)")
// 	if err != nil {
// 		fmt.Println("hello2")
// 		panic(err)
// 	}
// 	json_specification, err := json.Marshal(newproduct.Specification)
// 	_, err = insertStatement.Exec(newproduct.Product_id, newproduct.Name, newproduct.Sku, newproduct.Category_id, newproduct.Price, json_specification)
// 	if err != nil {
// 		fmt.Println("hello2")
// 		panic(err)
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(newproduct)
// }

// func AddCategory(w http.ResponseWriter, r *http.Request) {
// 	var category Category_master
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
// 	}

// 	err = json.Unmarshal(reqBody, &category)
// 	if err != nil {
// 		fmt.Fprintf(w, "error unmarshalling")
// 	}
// 	fmt.Printf("%+v", category)
// 	insertStatement, err := db.Prepare("INSERT INTO category_master(category_id,category_name) VALUES($1,$2)")
// 	if err != nil {
// 		fmt.Println("hello2")
// 		panic(err)
// 	}
// 	_, err = insertStatement.Exec(category.Category_id, category.Category_name)
// 	if err != nil {
// 		fmt.Println("hello2")
// 		panic(err)
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(category)
// }

// func GetCategory(w http.ResponseWriter, r *http.Request) {
// 	getCategory_id := mux.Vars(r)["id"]
// 	fmt.Println(getCategory_id)
// 	if CheckCategory_id(getCategory_id) == false {
// 		result := fmt.Sprintf("The product does not exsits enter a valid product id")
// 		json.NewEncoder(w).Encode(result)
// 		return
// 	}
// 	rows, err := db.Query("SELECT * from category_master  WHERE category_id = $1", getCategory_id)
// 	if err != nil {
// 		fmt.Println("err in selecting category")
// 	}
// 	defer rows.Close()
// 	var category Category_master

// 	for rows.Next() {

// 		err := rows.Scan(&category.Category_id, &category.Category_name)
// 		if err != nil {
// 			fmt.Println("error scaning")
// 			log.Fatal(err)
// 		}

// 		json.NewEncoder(w).Encode(category)
// 	}
// }

// func GetAllCategory(w http.ResponseWriter, r *http.Request) {
// 	x := mux.Vars(r)["id"]
// 	AllCategory := []Category_master{}
// 	page_no, err := strconv.Atoi(x)
// 	if err != nil {
// 		fmt.Println("page no invalid")
// 	}
// 	endlimit := page_no * 20
// 	startlimit := endlimit - 20
// 	fmt.Println(startlimit)
// 	rows, err := db.Query("Select * from category_master ")
// 	defer rows.Close()
// 	var category Category_master
// 	// var rawContent string
// 	catgorylist := []map[string]any{}

// 	for rows.Next() {
// 		err := rows.Scan(&category.Category_id, &category.Category_name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// err = json.Unmarshal([]byte(rawContent), &product.Specification)
// 		// if err != nil {
// 		// 	fmt.Println("error marshaling")
// 		// }

// 		AllCategory = append(AllCategory, category)

// 		// result := fmt.Sprintln(product.Price, product.Name)
// 		// json.NewEncoder(w).Encode(result)
// 	}
// 	endlimit = int(math.Min(float64(len(AllCategory)), float64(endlimit)))
// 	if startlimit >= 0 && startlimit < endlimit {
// 		fmt.Println(AllCategory[startlimit:endlimit])
// 	}
// 	for _, v := range AllCategory {
// 		newcategory := map[string]any{
// 			" id":  v.Category_id,
// 			"name": v.Category_name,
// 		}
// 		catgorylist = append(catgorylist, newcategory)
// 	}
// 	json.NewEncoder(w).Encode(catgorylist)

// }

// func DeleteCategory(w http.ResponseWriter, r *http.Request) {
// 	x := mux.Vars(r)["id"]

// 	res, err := db.Exec("DELETE FROM category_master WHERE category_id=$1", x)

// 	if err == nil {

// 		count, err := res.RowsAffected()
// 		if err == nil {
// 			if count == 0 {
// 				result := fmt.Sprint("The value category_id Does not exsits, eneter valid id")
// 				json.NewEncoder(w).Encode(result)
// 				return
// 			}
// 			result := fmt.Sprint("The value is deleted sucessfully")

// 			json.NewEncoder(w).Encode(result)
// 		}

// 	}

// 	return

// }

// func UpdateCategory(w http.ResponseWriter, r *http.Request) {
// 	var category Category_master
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data withid and name only in order to update")
// 	}
// 	fmt.Println(reqBody)
// 	err = json.Unmarshal(reqBody, &category)
// 	if err != nil {
// 		fmt.Printf("error marshaling")
// 	}
// 	fmt.Println("abababa")
// 	fmt.Println(category)
// 	if CheckCategory_id(category.Category_id) == false {
// 		result := fmt.Sprintf("The category does not exsits enter a valid category id")
// 		json.NewEncoder(w).Encode(result)
// 		return
// 	}
// 	fmt.Println("hilli")
// 	rows, err := db.Query("SELECT * from category_master WHERE category_id = $1", category.Category_id)
// 	if err != nil {
// 		fmt.Println("err in selecting category")
// 	}
// 	defer rows.Close()
// 	var exsisting_category Category_master
// 	for rows.Next() {
// 		fmt.Println("worki")
// 		err := rows.Scan(&exsisting_category.Category_id, &exsisting_category.Category_name)
// 		if err != nil {
// 			fmt.Println("error scaning")
// 			log.Fatal(err)
// 		}

// 		if category.Category_name == "" {
// 			category.Category_name = exsisting_category.Category_name
// 		}

// 		fmt.Println(category)
// 		// db.Query("UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;", newproduct.Name, newproduct.Sku, newproduct.Price, json_specification, newproduct.Product_id)
// 		db.Query("UPDATE category_master SET category_name=$1 WHERE category_id =$1;", category.Category_name, category.Category_id)
// 		if err != nil {
// 			fmt.Println("")
// 		}
// 	}
// }

// // check product id
// func CheckProduct_id(id string) bool {

// 	rows, err := db.Query("SELECT * from product_master WHERE product_id = $1", id)
// 	fmt.Println(id)

// 	if rows.Next() == false {
// 		return false
// 	}
// 	if err != nil {
// 		fmt.Println(err)
// 		return false
// 	}
// 	defer rows.Close()
// 	return true
// }

// // checkCategory_id
// func CheckCategory_id(id string) bool {
// 	fmt.Println(id)

// 	rows, err := db.Query("SELECT * from category_master WHERE category_id = $1", id)
// 	fmt.Println("check1")
// 	if err == nil {
// 		// fmt.Println(err)
// 		// return false
// 	}

// 	if rows.Next() == false {
// 		return false
// 	}
// 	defer rows.Close()
// 	return true
// }

// func UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	var newproduct Product_master
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
// 	}

// 	err = json.Unmarshal(reqBody, &newproduct)
// 	fmt.Println(newproduct)
// 	if CheckProduct_id(newproduct.Product_id) == false {
// 		result := fmt.Sprintf("The product does not exsits enter a valid product id")
// 		json.NewEncoder(w).Encode(result)
// 		return
// 	}
// 	fmt.Println("hilli")
// 	rows, err := db.Query("SELECT * from product_master WHERE product_id = $1", newproduct.Product_id)
// 	if err != nil {
// 		fmt.Println("err in selecting product")
// 	}
// 	defer rows.Close()
// 	var exsisting_product Product_master
// 	var rawContent []byte
// 	for rows.Next() {
// 		fmt.Println("worki")
// 		err := rows.Scan(&exsisting_product.Product_id, &exsisting_product.Name, &exsisting_product.Sku, &exsisting_product.Category_id, &exsisting_product.Price, &rawContent)
// 		if err != nil {
// 			fmt.Println("error scaning")
// 			log.Fatal(err)
// 		}

// 		err = json.Unmarshal(rawContent, &exsisting_product.Specification)
// 		if err != nil {
// 			fmt.Println("error unmarshaling")
// 		}
// 		if newproduct.Name == "" {
// 			newproduct.Name = exsisting_product.Name
// 		}
// 		if newproduct.Price == 0 {
// 			newproduct.Price = exsisting_product.Price
// 		}
// 		if newproduct.Sku == "" {
// 			newproduct.Sku = exsisting_product.Sku
// 		}
// 		if newproduct.Specification == nil {
// 			newproduct.Specification = exsisting_product.Specification
// 		}
// 		if newproduct.Category_id == "" {
// 			newproduct.Category_id = exsisting_product.Category_id
// 		}
// 		if newproduct.Category_id != exsisting_product.Category_id {
// 			result := fmt.Sprintf("Cannot alter the category id for product!!!\n please update in category table")
// 			json.NewEncoder(w).Encode(result)
// 			return
// 		}

// 		json_specification, err := json.Marshal(newproduct.Specification)

// 		fmt.Println(newproduct)
// 		// db.Query("UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;", newproduct.Name, newproduct.Sku, newproduct.Price, json_specification, newproduct.Product_id)
// 		db.Query("UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;", newproduct.Name, newproduct.Sku, newproduct.Price, json_specification, newproduct.Product_id)
// 		if err != nil {
// 			fmt.Println("")
// 		}
// 	}
// }

// func DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	x := mux.Vars(r)["id"]

// 	res, err := db.Exec("DELETE FROM product_master WHERE product_id=$1", x)

// 	if err == nil {

// 		count, err := res.RowsAffected()
// 		if err == nil {
// 			if count == 0 {
// 				result := fmt.Sprint("The value Product_Id Does not exsits, eneter valid id")
// 				json.NewEncoder(w).Encode(result)
// 				return
// 			}
// 			result := fmt.Sprint("The value is deleted sucessfully")
// 			json.NewEncoder(w).Encode(result)
// 		}

// 	}

// 	return

// }

// func GetProduct(w http.ResponseWriter, r *http.Request) {
// 	getproduct_id := mux.Vars(r)["id"]
// 	fmt.Println(getproduct_id)
// 	if CheckProduct_id(getproduct_id) == false {
// 		result := fmt.Sprintf("The product does not exsits enter a valid product id")
// 		json.NewEncoder(w).Encode(result)
// 		return
// 	}

// 	rows, err := db.Query("SELECT product_master.product_id,product_master.name,product_master.sku,product_master.price,category_master.category_name,product_master.specification  from product_master JOIN category_master ON product_master.category_id = category_master.category_id WHERE product_id = $1", getproduct_id)
// 	if err != nil {
// 		fmt.Println("err in selecting product")
// 	}
// 	defer rows.Close()
// 	var product Product_master
// 	var rawContent string
// 	var cateory_name string
// 	for rows.Next() {

// 		err := rows.Scan(&product.Product_id, &product.Name, &product.Sku, &product.Price, &cateory_name, &rawContent)
// 		if err != nil {
// 			fmt.Println("error scaning")
// 			log.Fatal(err)
// 		}
// 		err = json.Unmarshal([]byte(rawContent), &product.Specification)
// 		if err != nil {
// 			fmt.Println("error unmarshaling")
// 		}
// 		result := fmt.Sprint("details of product", product.Product_id, product.Name, product.Sku, product.Price, cateory_name, &product.Specification)
// 		json.NewEncoder(w).Encode(result)
// 	}
// }

// func GetAllProducts(w http.ResponseWriter, r *http.Request) {
// 	x := mux.Vars(r)["id"]
// 	allproducts := []Product_master{}
// 	page_no, err := strconv.Atoi(x)
// 	if err != nil {
// 		fmt.Println("page no invalid")
// 	}
// 	endlimit := page_no * 20
// 	startlimit := endlimit - 20
// 	fmt.Println(startlimit)
// 	rows, err := db.Query("Select product_id,name ,price  from product_master ")
// 	defer rows.Close()
// 	var product Product_master
// 	// var rawContent string
// 	productlist := []map[string]any{}

// 	for rows.Next() {
// 		err := rows.Scan(&product.Product_id, &product.Name, &product.Price)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// err = json.Unmarshal([]byte(rawContent), &product.Specification)
// 		// if err != nil {
// 		// 	fmt.Println("error marshaling")
// 		// }

// 		allproducts = append(allproducts, product)

// 		// result := fmt.Sprintln(product.Price, product.Name)
// 		// json.NewEncoder(w).Encode(result)
// 	}
// 	endlimit = int(math.Min(float64(len(allproducts)), float64(endlimit)))
// 	if startlimit >= 0 && startlimit < endlimit {
// 		fmt.Println(allproducts[startlimit:endlimit])
// 	}
// 	for _, v := range allproducts {
// 		newprod := map[string]any{
// 			"product_id": v.Product_id,
// 			"name":       v.Name,
// 			"price":      v.Price,
// 		}
// 		productlist = append(productlist, newprod)
// 	}
// 	json.NewEncoder(w).Encode(productlist)

// }

// func AddInventory(w http.ResponseWriter, r *http.Request) {
// 	var inventory Inventory
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
// 	}

// 	err = json.Unmarshal(reqBody, &inventory)
// 	if err != nil {
// 		fmt.Fprintf(w, "error unmarshalling")
// 	}
// 	fmt.Printf("%+v", inventory)
// 	if CheckProduct_id(inventory.Product_id) == false {
// 		result := fmt.Sprintf("The product does not exsits enter a valid product id")
// 		json.NewEncoder(w).Encode(result)
// 		return
// 	}
// 	insertStatement, err := db.Prepare("INSERT INTO inventory(product_id,quantity) VALUES($1,$2)")
// 	if err != nil {
// 		fmt.Println("hello2")
// 		panic(err)
// 	}
// 	_, err = insertStatement.Exec(inventory.Product_id, inventory.Quantity)
// 	if err != nil {
// 		fmt.Println("hello2")
// 		panic(err)
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(inventory)
// }

// func UpdateInventory(w http.ResponseWriter, r *http.Request) {
// 	var item Inventory
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the product_id and name only in order to insert")
// 	}

// 	err = json.Unmarshal(reqBody, &item)
// 	fmt.Println(item)
// 	if CheckProduct_id(item.Product_id) == false {
// 		result := fmt.Sprintf("The product does not exsits enter a valid product id")
// 		json.NewEncoder(w).Encode(result)
// 		return
// 	}
// 	fmt.Println("hilli")
// 	rows, err := db.Query("SELECT * from inventory WHERE product_id = $1", item.Product_id)
// 	if err != nil {
// 		fmt.Println("err in selecting product")
// 	}
// 	defer rows.Close()
// 	var exsisting_product Inventory
// 	for rows.Next() {
// 		fmt.Println("worki")
// 		err := rows.Scan(&exsisting_product.Product_id, &exsisting_product.Quantity)
// 		if err != nil {
// 			fmt.Println("error scaning")
// 			log.Fatal(err)
// 		}
// 		if item.Quantity <= 0 {
// 			item.Quantity = 0
// 		}

// 		fmt.Println(item)
// 		// db.Query("UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5;", newproduct.Name, newproduct.Sku, newproduct.Price, json_specification, newproduct.Product_id)
// 		db.Query("UPDATE inventory SET quantity=$1 WHERE product_id =$2", item.Quantity, item.Product_id)
// 		if err != nil {
// 			fmt.Println("")
// 		}
// 	}
// }
// func GetInventory(w http.ResponseWriter, r *http.Request) {
// 	getproduct_id := mux.Vars(r)["id"]
// 	fmt.Println(getproduct_id)
// 	if CheckProduct_id(getproduct_id) == false {
// 		result := fmt.Sprintf("The product does not exsits enter a valid product id")
// 		json.NewEncoder(w).Encode(result)
// 		return
// 	}

// 	rows, err := db.Query("SELECT * from inventory WHERE product_id = $1", getproduct_id)
// 	if err != nil {
// 		fmt.Println("err in selecting product")
// 	}
// 	defer rows.Close()
// 	var item Inventory

// 	for rows.Next() {

// 		err := rows.Scan(&item.Product_id, &item.Quantity)
// 		if err != nil {
// 			fmt.Println("error scaning")
// 			log.Fatal(err)
// 		}
// 		// result := fmt.Sprint("Inventory detais of prodct", product.Product_id, product.Name, product.Sku, product.Price, cateory_name, &product.Specification)
// 		json.NewEncoder(w).Encode(item)
// 	}
// }

// func DeleteInventory(w http.ResponseWriter, r *http.Request) {
// 	x := mux.Vars(r)["id"]

// 	res, err := db.Exec("DELETE FROM inventory WHERE product_id=$1", x)

// 	if err == nil {

// 		count, err := res.RowsAffected()
// 		if err == nil {
// 			if count == 0 {
// 				result := fmt.Sprint("The value Product_Id Does not exsits, eneter valid id")
// 				json.NewEncoder(w).Encode(result)
// 				return
// 			}
// 			result := fmt.Sprint("The value is deleted sucessfully")
// 			json.NewEncoder(w).Encode(result)
// 		}

// 	}

// 	return

// }

// func GetAllInventory(w http.ResponseWriter, r *http.Request) {
// 	x := mux.Vars(r)["id"]
// 	allproducts := []Inventory{}
// 	page_no, err := strconv.Atoi(x)
// 	if err != nil {
// 		fmt.Println("page ni invalid")
// 	}
// 	endlimit := page_no * 20
// 	startlimit := endlimit - 20
// 	fmt.Println(startlimit)
// 	rows, err := db.Query("Select *  from inventory ")
// 	defer rows.Close()
// 	var product Inventory
// 	// var rawContent string
// 	productlist := []map[string]any{}

// 	for rows.Next() {
// 		err := rows.Scan(&product.Product_id, &product.Quantity)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// err = json.Unmarshal([]byte(rawContent), &product.Specification)
// 		// if err != nil {
// 		// 	fmt.Println("error marshaling")
// 		// }

// 		allproducts = append(allproducts, product)

// 		// result := fmt.Sprintln(product.Price, product.Name)
// 		// json.NewEncoder(w).Encode(result)
// 	}
// 	endlimit = int(math.Min(float64(len(allproducts)), float64(endlimit)))
// 	if startlimit >= 0 && startlimit < endlimit {
// 		fmt.Println(allproducts[startlimit:endlimit])
// 	}
// 	for _, v := range allproducts {
// 		newprod := map[string]any{
// 			"Product_id": v.Product_id,
// 			"Quantity":   v.Quantity,
// 		}
// 		productlist = append(productlist, newprod)
// 	}
// 	json.NewEncoder(w).Encode(allproducts)

// }
// func CreateCart(w http.ResponseWriter, r *http.Request) {
// 	name := mux.Vars(r)["name"]
// 	fmt.Println(name)

// 	id := uuid.New()
// 	fmt.Println("Generated UUID:")

// 	var newcart Cart
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
// 	}

// 	err = json.Unmarshal(reqBody, &newcart)
// 	if err != nil {
// 		fmt.Fprintf(w, "error unmarshalling")
// 	}
// 	insertStatement, err := db.Prepare("INSERT INTO cart_reference(reference_id,create_date,username) VALUES($1,now(),$2)")
// 	if err != nil {
// 		fmt.Println("hello2")
// 		panic(err)
// 	}
// 	_, err = insertStatement.Exec(id, name)
// 	if err != nil {
// 		fmt.Println("hello2")
// 		panic(err)
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(id)

// }

// func CheckQuantity(Product_id string, Quantity int) bool {
// 	rows, err := db.Query("SELECT quantity from inventory WHERE product_id = $1", Product_id)
// 	if err != nil {
// 		fmt.Println("errr", err)
// 	}
// 	fmt.Println(Product_id, Quantity)
// 	var w http.ResponseWriter
// 	for rows.Next() {
// 		var stock int

// 		rows.Scan(&stock)
// 		fmt.Println(stock, "stock")
// 		if stock == 0 {
// 			fmt.Println("h1")
// 			result := fmt.Sprintf("OUT OF STOCK!!!SOLD OUT")
// 			return false
// 			json.NewEncoder(w).Encode(result)

// 		}
// 		if stock < Quantity {
// 			fmt.Println("h2")
// 			result := fmt.Sprintf("The selected quantity is not available please choose less number of items")
// 			return false
// 			json.NewEncoder(w).Encode(result)

// 		} else {
// 			fmt.Println("h3")
// 			cartcount := stock - Quantity
// 			db.Exec("UPDATE inventory SET quantity=$1 WHERE product_id =$2", cartcount, Product_id)
// 			return true
// 		}
// 	}

// 	defer rows.Close()
// 	return true

// }

// func CheckCartProduct(Reference_id string, Product_id string, Quantity int) bool {
// 	rows, err := db.Query("SELECT quantity from cart WHERE product_id = $1 and reference_id= $2", Product_id, Reference_id)
// 	if err != nil {
// 		fmt.Println("errr", err)
// 	}
// 	fmt.Println(Reference_id, Product_id, Quantity)
// 	var w http.ResponseWriter
// 	for rows.Next() {
// 		var cartitem int
// 		rows.Scan(&cartitem)
// 		cartcount := Quantity + cartitem
// 		db.Exec("UPDATE cart SET quantity=$1 WHERE product_id =$2 and reference_id=$3", cartcount, Product_id, Reference_id)
// 		result := fmt.Sprintf("selected quantity of product added to your cart")
// 		json.NewEncoder(w).Encode(result)

// 		return true
// 	}

// 	defer rows.Close()
// 	return false

// }

// func AddToCart(w http.ResponseWriter, r *http.Request) {

// 	var newcart Cart
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
// 	}

// 	err = json.Unmarshal(reqBody, &newcart)
// 	if err != nil {
// 		fmt.Fprintf(w, "error unmarshalling")
// 	}
// 	//check if valid product is there or not
// 	if CheckProduct_id(newcart.Product_id) == false {
// 		result := fmt.Sprintf("The product does not exsits enter a valid product id")
// 		json.NewEncoder(w).Encode(result)
// 		return
// 	}

// 	// check if the amount is present inventory or not
// 	if CheckQuantity(newcart.Product_id, newcart.Quantity) == false {
// 		result := fmt.Sprintf("OUT OF STOCK!!!The selected quanitity is not available please select less number of items")
// 		json.NewEncoder(w).Encode(result)
// 		return
// 	}

// 	// update inventory
// 	// var item inventory
// 	// item.Product_id = newcart.Product_id
// 	// item.Quantity = newcart.Quantity

// 	// updateInventory(w, r)

// 	//if product not present then insert or else update
// 	if CheckCartProduct(newcart.Reference_id, newcart.Product_id, newcart.Quantity) == false {
// 		insertStatement, err := db.Prepare("INSERT INTO cart(reference_id,product_id,quantity) VALUES($1,$2,$3)")
// 		if err != nil {
// 			fmt.Println("hel")
// 			panic(err)
// 		}
// 		_, err = insertStatement.Exec(newcart.Reference_id, newcart.Product_id, newcart.Quantity)
// 		if err != nil {
// 			fmt.Println("hello2")
// 			panic(err)
// 		}
// 		result := fmt.Sprint("The product added successfully")
// 		json.NewEncoder(w).Encode(result)
// 		w.WriteHeader(http.StatusCreated)

// 	}

// 	//check if item is present in cart already then update

// 	// db.Query("UPDATE cart SET quantity=$1 WHERE product_id =$2 and reference_id=$3", newcart.Quantity, newcart.Product_id, newcart.Reference_id)

// 	// insertStatement, err := db.Prepare("INSERT INTO cart(reference_id,product_id,quantity) VALUES($1,$2,$3)")
// 	// if err != nil {
// 	// 	fmt.Println("hel")
// 	// 	panic(err)
// 	// }
// 	// _, err = insertStatement.Exec(newcart.Reference_id, newcart.Product_id, newcart.Quantity)
// 	// if err != nil {
// 	// 	fmt.Println("hello2")
// 	// 	panic(err)
// 	// }
// 	// w.WriteHeader(http.StatusCreated)
// 	// json.NewEncoder(w).Encode(inventory)

// }

// func DeleteCart(w http.ResponseWriter, r *http.Request) {
// 	x := mux.Vars(r)["id"]

// 	res, err := db.Exec("DELETE FROM cart WHERE reference_id=$1", x)

// 	if err == nil {

// 		count, err := res.RowsAffected()
// 		if err == nil {
// 			if count == 0 {
// 				result := fmt.Sprint("The reference_id Does not exsits, eneter valid id")
// 				json.NewEncoder(w).Encode(result)
// 				return
// 			}
// 			result := fmt.Sprint("The cart item is deleted sucessfully")
// 			json.NewEncoder(w).Encode(result)
// 		}

// 	}

// 	return

// }
// func GetCart(w http.ResponseWriter, r *http.Request) {
// 	x := mux.Vars(r)["id"]
// 	items := []Inventory{}
// 	// page_no, err := strconv.Atoi(x)
// 	// if err != nil {
// 	// 	fmt.Println("page no invalid")
// 	// }
// 	// // endlimit := page_no * 20
// 	// startlimit := endlimit - 20
// 	// fmt.Println(startlimit)
// 	rows, err := db.Query("Select product_id,quantity from cart where reference_id=$1 ", x)
// 	if err != nil {
// 		fmt.Println("page no invalid")
// 	}
// 	defer rows.Close()
// 	var item Inventory
// 	// itemlist := []map[string]any{}

// 	for rows.Next() {
// 		err := rows.Scan(&item.Product_id, &item.Quantity)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// err = json.Unmarshal([]byte(rawContent), &product.Specification)
// 		// if err != nil {
// 		// 	fmt.Println("error marshaling")
// 		// }

// 		items = append(items, item)

// 		// result := fmt.Sprintln(product.Price, product.Name)
// 		// json.NewEncoder(w).Encode(result)
// 	}
// 	// endlimit = int(math.Min(float64(len(AllCategory)), float64(endlimit)))
// 	// if startlimit >= 0 && startlimit < endlimit {
// 	// 	fmt.Println(AllCategory[startlimit:endlimit])
// 	// }
// 	// for _, v := range AllCategory {
// 	// 	newcategory := map[string]any{
// 	// 		" id":  v.Category_id,
// 	// 		"name": v.Category_name,
// 	// 	}
// 	// 	catgorylist = append(catgorylist, newcategory)
// 	// }

// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(items)

// }

func main() {
	Support.DbConnect()
	router := mux.NewRouter()
	// router.HandleFunc("/", HomeLink)
	//product handlers
	router.HandleFunc("/addProduct/", Product.AddProducts).Methods("POST")
	router.HandleFunc("/getProduct/{id}", Product.GetProduct).Methods("GET")
	router.HandleFunc("/getProducts/{id}", Product.GetAllProducts).Methods("GET")
	router.HandleFunc("/deleteProduct/{id}", Product.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/updateProduct", Product.UpdateProduct).Methods("PUT")
	// category handlers call
	router.HandleFunc("/addCategory", Category.AddCategory).Methods("POST")
	router.HandleFunc("/getCategory/{id}", Category.GetCategory).Methods("GET")
	router.HandleFunc("/getAllCategory/{id}", Category.GetAllCategory).Methods("GET")
	router.HandleFunc("/deleteCategory/{id}", Category.DeleteCategory).Methods("DELETE")
	router.HandleFunc("/updateCategory", Category.UpdateCategory).Methods("PUT")
	// inventory handlers call
	router.HandleFunc("/addInventory", Inventory.AddInventory).Methods("POST")
	router.HandleFunc("/getInventory/{id}", Inventory.GetInventory).Methods("GET")
	router.HandleFunc("/getAllInventory/{id}", Inventory.GetAllInventory).Methods("GET")
	router.HandleFunc("/deleteInventory/{id}", Inventory.DeleteInventory).Methods("DELETE")
	router.HandleFunc("/updateInventory", Inventory.UpdateInventory).Methods("PUT")

	//cart handlers

	router.HandleFunc("/addToCart", Cart.AddToCart).Methods("POST")
	router.HandleFunc("/addItemsToCart/{Ref_id}", Cart.AddItemsToCart).Methods("POST")
	router.HandleFunc("/createCart/{name}", Cart.CreateCart).Methods("POST")

	router.HandleFunc("/getCart/{id}", Cart.GetCart).Methods("GET")
	router.HandleFunc("/deleteCart/{id}", Cart.DeleteCart).Methods("DELETE")
	// router.HandleFunc("/updateCart", updateCart).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8083", router))
}
