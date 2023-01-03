package Support

import (
	"fmt"
)

// check product id
func CheckReference_id(id string) bool {

	rows, err := DB.Query("SELECT * from cart_reference WHERE reference_id = $1", id)
	fmt.Println(id)

	if !rows.Next() {
		return false
	}
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	return true
}
