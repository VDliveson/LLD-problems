package internal

type MenuItem struct {
	Id    int
	Name  string
	Price float64
}

type Order struct {
	Id         int
	Items      []MenuItem
	TableId    int
	CustomerId int
}

type Customer struct {
	Id     int
	Name   string
	Mobile string
}

type Table struct {
	Id       int
	Capacity int
}

type Bill struct {
	Id         int
	Items      []MenuItem
	Amount     float64
	CustomerId int
}
