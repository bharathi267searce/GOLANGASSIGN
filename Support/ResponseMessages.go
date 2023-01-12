package Support

// error responses
var InvalidReferenceId string = "The Reference Id is INVALID!!! Please enter a valid Reference Id "
var InvalidCategoryId string = "The Category Id is INVALID!!! Please enter a valid Category Id "
var InvalidProductId string = "The Product Id is INVALID!!! Please enter a valid Product Id "

var ErrorMarshaling string = " Error while Marshaling your data"
var ErrorUnMarshaling string = " Error while Unmarshaling your response"
var ErrorUnMarshalingSpecification string = " Error while UnMarshaling your Specification"
var ErrorMarshalingSpecification string = " Error while Marshaling your Specification"

// db statements
var PrepareStatementError string = "Error while preparing sql statement"
var ExecStatementError string = "Error while executing sql statement"

var InvalidIDType string = " The type of ID entered is invalid "
var InvalidProductFormat string = "Enter the Product details as \nProductID \nProductName \nSKU \nCategoryID \nPrice \n Specification"
var InvalidCategoryFormat string = "Enter the Product details as \nProductID \nProductName \nSKU \nCategoryID \nPrice \n Specification"
var InvalidCartFormat string = "Enter the Product details as \nProductID \nProductName \nSKU \nCategoryID \nPrice \n Specification"
var InvalidInventoryFormat string = "Enter the Product details as \nProductID \nProductName \nSKU \nCategoryID \nPrice \n Specification"
var InvalidFormatDisplay string = "Invalid format"

// already exsists
var ProductExsist string = "The Product already exsists"
var CategoryExsist string = "The Category already exsists"

var ReferenceIdExsist string = "The Reference already exsists"

// invalid page number
var InvalidPageNo string = "The page number is invalid,Please enter a valid number"

var ItemsPerPage string = "The items per page is invalid,Please enter a valid number"
var ErrorScaning string = "Unable to scan the data"

var ErrorGetData string = "Unable to get the data from database"
var ErrorCategoryId string = "Cannot alter the category id for product!!!\n please update in category table"

// success messages

var ProductInserted string = "The product has been inserted successfully"
var ProductDeleted string = "The product is deleted sucessfully"
var ProductUpdated string = " The product is Updated successfully"

var CategoryInserted string = "The Category has been inserted sucessfully"
var CategoryDeleted string = "The Category is deleted sucessfully"
var CategoryUpdated string = " The Category is Updated successfully"

var InventoryInserted string = "The Inventory has been inserted sucessfully"
var InventoryDeleted string = "The Inventory is deleted sucessfully"
var InventoryUpdated string = " The Inventory is Updated successfully"
var ErrorScaningInput string = "Error while Scaning input "
var ItemExsistInventory string = "Item already Exsist ,Update inventory"

var OutOfStock string = "OUT OF STOCK!!!SOLD OUT"
var UnvailableQuanity string = "The selected quantity is not available please choose less number of items"
var NoQuantityCheck string = "Unable to check the quantity"
var UpdateCartItem string = "Updated the exsisting item in Cart"
var InsertedCartItem string = "Inserted cart Item sucessfully"
var QuantityMustBePositive string = "Quantity must be positive great than zero"
var CartDeleted string = "The product has been removed"
