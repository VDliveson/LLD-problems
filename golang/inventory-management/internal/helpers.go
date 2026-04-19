package inventory_management

import "fmt"

func DisplayProducts(products []*Product) {
	for _, product := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", product.ID, product.Name, product.Price)
	}
}

func DisplayWarehouses(warehouses []*Warehouse) {
	for _, warehouse := range warehouses {
		fmt.Printf("Warehouse ID: %d\n", warehouse.ID)
		for productID, quantity := range warehouse.Items {
			fmt.Printf("  Product ID: %d, Quantity: %d\n", productID, quantity)
		}
	}
}
