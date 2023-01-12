package query

// product queries
var AddProduct string = "INSERT INTO product_master(product_id,name,sku,category_id,price,specification) VALUES($1,$2,$3,$4,$5,$6)"
var DeleteProduct string = "DELETE FROM product_master WHERE product_id=$1"
var GetAllproduct string = "Select product_id,name ,price  from product_master"
var GetProduct string = "SELECT product_master.product_id,product_master.name,product_master.sku,product_master.price,category_master.category_name,product_master.specification  from product_master JOIN category_master ON product_master.category_id = category_master.category_id WHERE product_id = $1"
var GetProductInfo string = "SELECT *  from product_master where product_id = $1"
var UpdateProduct string = "UPDATE product_master SET name=$1,sku=$2, price=$3,specification=$4 WHERE product_id =$5"

// cart queries
var AddCart string = "INSERT INTO cart(reference_id,product_id,quantity) VALUES($1,$2,$3)"
var UpdateCartItem string = "UPDATE cart SET quantity=$1 WHERE product_id =$2 and reference_id=$3"
var CheckCartItem string = "SELECT quantity from cart WHERE product_id = $1 and reference_id= $2"
var CreateCart string = "INSERT INTO cart_reference(reference_id,create_date,username) VALUES($1,now(),$2)"
var GetCart string = "select d.product_id,p.price,c.category_name,d.quantity from product_master p  inner join  category_master c on c.category_id=p.category_id inner join cart d on d.product_id = p.product_id where d.reference_id =$1 "
var GetReferenceId string = "SELECT * from cart_reference WHERE reference_id = $1"
var DeleteCart string = "DELETE from cart where reference_id=$1 and product_id=$2"

// category queries
var AddCategory string = "INSERT INTO category_master(category_id,category_name) VALUES($1,$2)"
var DeleteCategory string = "DELETE FROM category_master WHERE category_id=$1"
var GetAllCategory string = "Select * from category_master "
var GetCategory string = "SELECT * from category_master  WHERE category_id = $1"
var UpdateCategory string = "UPDATE category_master SET category_name=$1 WHERE category_id =$2;"

// inventory queries
var AddInventory string = "INSERT INTO inventory(product_id,quantity) VALUES($1,$2)"
var DeleteInventory string = "DELETE FROM inventory WHERE product_id=$1"
var GetAllInventory string = "Select *  from inventory "
var GetInventory string = "SELECT * from inventory WHERE product_id = $1"
var UpdateInventory string = "UPDATE inventory SET quantity=$1 WHERE product_id =$2"
var CheckQuantity string = "SELECT quantity from inventory WHERE product_id = $1"
