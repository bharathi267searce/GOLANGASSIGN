package Support

import (
	"fmt"
)

// check product id
func CheckProduct_id(id string) bool {

	rows, err := DB.Query("SELECT * from product_master WHERE product_id = $1", id)
	fmt.Println(id)

	if !rows.Next() {
		return true
	}
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	return true
}
