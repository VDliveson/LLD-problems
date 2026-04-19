package inventory_management

import (
	"fmt"
	"sync"
)

type InventoryService struct {
	Products       map[int]*Product
	productMutex   sync.RWMutex
	warehouses     map[int]*Warehouse
	warehouseMutex sync.RWMutex
}

func NewInventoryService() *InventoryService {
	return &InventoryService{
		Products:   make(map[int]*Product),
		warehouses: make(map[int]*Warehouse),
	}
}

func (s *InventoryService) AddProduct(ID int, name string, price float64) error {
	s.productMutex.Lock()
	defer s.productMutex.Unlock()
	if _, exists := s.Products[ID]; exists {
		return fmt.Errorf("product with ID %d already exists", ID)
	}

	product := &Product{
		ID:    ID,
		Name:  name,
		Price: price,
	}
	s.Products[product.ID] = product
	return nil
}

func (s *InventoryService) CreateWarehouse(ID int) error {
	s.warehouseMutex.Lock()
	defer s.warehouseMutex.Unlock()

	if _, exists := s.warehouses[ID]; exists {
		return fmt.Errorf("warehouse with ID %d already exists", ID)
	}

	inventory := &Warehouse{
		ID:    ID,
		Items: make(map[int]int),
	}
	s.warehouses[inventory.ID] = inventory
	return nil
}

func (s *InventoryService) UpdateWarehouseInventory(ID int, productID int, quantity int, add bool) error {
	s.productMutex.RLock()
	if _, exists := s.Products[productID]; !exists {
		s.productMutex.RUnlock()
		return fmt.Errorf("product with ID %d does not exist", productID)
	}
	s.productMutex.RUnlock()

	s.warehouseMutex.RLock()
	inventory, exists := s.warehouses[ID]
	if !exists {
		s.warehouseMutex.RUnlock()
		return fmt.Errorf("warehouse with ID %d does not exist", ID)
	}
	s.warehouseMutex.RUnlock()

	inventory.Mu.Lock()
	defer inventory.Mu.Unlock()

	if add {
		inventory.Items[productID] += quantity
	} else {
		if currentQty, exists := inventory.Items[productID]; exists {
			if currentQty < quantity {
				quantity = currentQty
			}
			inventory.Items[productID] -= quantity
		}
	}
	return nil
}

func (s *InventoryService) ListProducts() []*Product {
	s.productMutex.RLock()
	defer s.productMutex.RUnlock()

	products := make([]*Product, 0, len(s.Products))
	for _, product := range s.Products {
		products = append(products, product)
	}
	return products
}

func (s *InventoryService) GetStock(warehouseID int, productID int) (int, error) {
	s.warehouseMutex.RLock()
	inventory, exists := s.warehouses[warehouseID]
	s.warehouseMutex.RUnlock()
	if !exists {
		return 0, fmt.Errorf("warehouse with ID %d does not exist", warehouseID)
	}

	inventory.Mu.RLock()
	defer inventory.Mu.RUnlock()

	quantity, exists := inventory.Items[productID]
	if !exists {
		return 0, fmt.Errorf("product with ID %d not found in warehouse %d", productID, warehouseID)
	}
	return quantity, nil
}

func (s *InventoryService) ListWarehouses() []*Warehouse {
	s.warehouseMutex.RLock()
	defer s.warehouseMutex.RUnlock()

	warehouses := make([]*Warehouse, 0, len(s.warehouses))
	for _, warehouse := range s.warehouses {
		warehouses = append(warehouses, warehouse)
	}
	return warehouses
}
