package product

import (
	"fmt"
)

func Get_product() string {
	fmt.Println("Enter the name of the product: ")
	var search_product string
	fmt.Scanln(&search_product)
	return search_product
}
