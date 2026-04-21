package internal

import "sync"

type User struct {
	Id    int
	Name  string
	Phone string
	Email string
}

type Group struct {
	Id      int
	Name    string
	Members map[int]struct{}
	mu      sync.Mutex
}
type Expense struct {
	Id          int
	Description string
	Amount      float64
	PaidBy      int
	SplitType   string
	Split       []Split
	mu          sync.Mutex
}

type SplitService struct {
	balances map[int]map[int]float64
	groups   map[int]*Group
	expenses map[int]*Expense
	users    map[int]*User
	mu       sync.RWMutex
}

type Split struct {
	UserId int
	Amount float64
}

type Balance struct {
	FromUserId int
	ToUserId   int
	Amount     float64
}
