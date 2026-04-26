package internal

import "sync"

type User struct {
	id    int
	name  string
	email string
}

type Wallet struct {
	id     int
	userId int
	amount float64
	mu     sync.RWMutex
}

type Transaction struct {
	Id         int
	FromUserId int
	ToUserId   int
	Amount     float64
}
