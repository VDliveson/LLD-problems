package inventory_management

import "sync"

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Warehouse struct {
	ID    int
	Items map[int]int
	Mu    sync.RWMutex
}
