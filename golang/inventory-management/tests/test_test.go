package tests

import (
	internal "inventory-management/internal"
	"sync"
	"testing"
)

func TestInventoryService_Functional(t *testing.T) {
	s := internal.NewInventoryService()
	err := s.AddProduct(1, "Laptop", 1000.0)
	if err != nil {
		t.Fatal(err)
	}
	err = s.CreateWarehouse(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("AddStock", func(t *testing.T) {
		err := s.UpdateWarehouseInventory(1, 1, 10, true)
		if err != nil {
			t.Fatal(err)
		}
		val, _ := s.GetStock(1, 1)
		if val != 10 {
			t.Errorf("expected 10, got %d", val)
		}
	})

	t.Run("DeductStock", func(t *testing.T) {
		err := s.UpdateWarehouseInventory(1, 1, 4, false)
		if err != nil {
			t.Fatal(err)
		}
		val, _ := s.GetStock(1, 1)

		if val != 6 {
			t.Errorf("expected 6, got %d", val)
		}
	})

	t.Run("InsufficientStock", func(t *testing.T) {
		err := s.UpdateWarehouseInventory(1, 1, 10, false)
		if err != nil {
			t.Fatal(err)
		}
		val, _ := s.GetStock(1, 1)
		if val != 0 {
			t.Errorf("expected 0, got %d", val)
		}
	})
}

func TestInventoryService_Concurrency(t *testing.T) {
	s := internal.NewInventoryService()
	err := s.AddProduct(1, "Concurrency-Test", 1.0)
	if err != nil {
		t.Fatal(err)
	}
	err = s.CreateWarehouse(1)
	if err != nil {
		t.Fatal(err)
	}

	const workers = 100
	const incrementsPerWorker = 10
	var wg sync.WaitGroup

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerWorker; j++ {
				err := s.UpdateWarehouseInventory(1, 1, 1, true)
				if err != nil {
					t.Fatal(err)
				}
			}
		}()
	}

	wg.Wait()

	finalStock, _ := s.GetStock(1, 1)
	expected := workers * incrementsPerWorker
	if finalStock != expected {
		t.Errorf("race condition detected: expected %d, got %d", expected, finalStock)
	}
}

func TestInventoryService_NotFound(t *testing.T) {
	s := internal.NewInventoryService()

	err := s.UpdateWarehouseInventory(99, 1, 10, true)
	if err == nil {
		t.Error("expected error for non-existent warehouse")
	}

	err = s.CreateWarehouse(1)
	if err != nil {
		t.Fatal(err)
	}
	err = s.UpdateWarehouseInventory(1, 99, 10, true)
	if err == nil {
		t.Error("expected error for non-existent product")
	}
}
