package internal

import "sync"

type RestaurantService struct {
	MenuItems        map[int]MenuItem
	Orders           map[int]Order
	Customers        map[int]Customer
	Tables           map[int]Table
	Bills            map[int]Bill
	TableStatus      map[int]bool
	TableOrders      map[int][]int
	OrderMutex       sync.RWMutex
	MenuMutex        sync.RWMutex
	TableMutex       sync.RWMutex
	TableOrderMutex  sync.RWMutex
	TableStatusMutex sync.RWMutex
	CustomerMutex    sync.RWMutex
}

func NewRestaurantService() *RestaurantService {
	return &RestaurantService{
		MenuItems:   make(map[int]MenuItem),
		Orders:      make(map[int]Order),
		Customers:   make(map[int]Customer),
		Tables:      make(map[int]Table),
		Bills:       make(map[int]Bill),
		TableOrders: make(map[int][]int),
		TableStatus: make(map[int]bool),
	}
}

func (s *RestaurantService) AddMenuItem(name string, price float64) {
	s.MenuMutex.Lock()
	defer s.MenuMutex.Unlock()

	item := MenuItem{
		Id:    len(s.MenuItems) + 1,
		Name:  name,
		Price: price,
	}
	s.MenuItems[item.Id] = item
}

func (s *RestaurantService) PlaceOrder(items []MenuItem, tableId int, customerId int) {
	s.OrderMutex.Lock()
	defer s.OrderMutex.Unlock()

	s.TableOrderMutex.Lock()
	defer s.TableOrderMutex.Unlock()

	order := Order{
		Id:         len(s.Orders) + 1,
		Items:      items,
		TableId:    tableId,
		CustomerId: customerId,
	}
	s.Orders[order.Id] = order
	s.TableOrders[tableId] = append(s.TableOrders[tableId], order.Id)
}

func (s *RestaurantService) AddCustomer(name string, mobile string) {
	s.CustomerMutex.Lock()
	defer s.CustomerMutex.Unlock()

	customer := Customer{
		Id:     len(s.Customers) + 1,
		Name:   name,
		Mobile: mobile,
	}
	s.Customers[customer.Id] = customer
}

func (s *RestaurantService) AddTable(capacity int) {
	s.TableMutex.Lock()
	defer s.TableMutex.Unlock()

	table := Table{
		Id:       len(s.Tables) + 1,
		Capacity: capacity,
	}
	s.Tables[table.Id] = table
	s.TableStatus[table.Id] = false
}

func (s *RestaurantService) GenerateBill(tableId int, customerId int) *Bill {
	s.OrderMutex.Lock()
	defer s.OrderMutex.Unlock()

	orders := s.TableOrders[tableId]
	if len(orders) == 0 {
		return nil
	}

	totalAmount := 0.0
	bill := Bill{
		Id:         len(s.Bills) + 1,
		Items:      []MenuItem{},
		CustomerId: customerId,
	}

	itemCount := make(map[int]int)
	for _, orderId := range orders {
		order, ok := s.Orders[orderId]
		if !ok {
			continue
		}
		items := order.Items
		for _, item := range items {
			itemCount[item.Id]++
		}
	}

	for itemId, count := range itemCount {
		item, ok := s.MenuItems[itemId]
		if !ok {
			continue
		}
		bill.Items = append(bill.Items, item)
		totalAmount += item.Price * float64(count)
	}
	bill.Amount = totalAmount

	s.TableOrderMutex.Lock()
	defer s.TableOrderMutex.Unlock()

	s.TableStatusMutex.Lock()
	defer s.TableStatusMutex.Unlock()

	s.TableOrders[tableId] = []int{}
	s.TableStatus[tableId] = false
	s.Bills[bill.Id] = bill

	return &bill
}
