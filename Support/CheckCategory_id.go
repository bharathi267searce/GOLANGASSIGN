package Support

import (
	"fmt"
)

// checkCategory_id
func CheckCategory_id(id string) bool {
	fmt.Println(id)

	rows, err := DB.Query("SELECT * from category_master WHERE category_id = $1", id)
	fmt.Println("check1")
	if err == nil {
		// fmt.Println(err)
		// return false
	}

	if !rows.Next() {
		return false
	}
	defer rows.Close()
	return true
}
