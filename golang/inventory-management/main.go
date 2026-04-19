package main

import (
	internal "inventory-management/internal"
	"log"
)

func main() {
	inventoryService := internal.NewInventoryService()
	err := inventoryService.AddProduct(1, "Laptop", 999.99)
	if err != nil {
		log.Print(err)
	}

	err = inventoryService.AddProduct(2, "Smartphone", 499.99)
	if err != nil {
		log.Print(err)
	}

	err = inventoryService.CreateWarehouse(1)
	if err != nil {
		log.Print(err)
	}
	err = inventoryService.CreateWarehouse(2)
	if err != nil {
		log.Print(err)
	}

	err = inventoryService.UpdateWarehouseInventory(1, 1, 10, true) // Add 10 Laptops to Warehouse 1
	if err != nil {
		log.Print(err)
	}
	err = inventoryService.UpdateWarehouseInventory(1, 2, 20, true) // Add 20 Smartphones to Warehouse 1
	if err != nil {
		log.Print(err)
	}
	err = inventoryService.UpdateWarehouseInventory(2, 1, 50, false) // Removes 50 Laptops from Warehouse 2
	if err != nil {
		log.Print(err)
	}

	internal.DisplayProducts(inventoryService.ListProducts())
	internal.DisplayWarehouses(inventoryService.ListWarehouses())
}
