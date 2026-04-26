package internal

import (
	"fmt"
	"sync"
)

type WalletService struct {
	Wallets      map[int]*Wallet
	Users        map[int]User
	Transactions map[int]Transaction
	mu           sync.RWMutex
}

func NewWalletService() *WalletService {
	return &WalletService{
		Wallets:      make(map[int]*Wallet),
		Users:        make(map[int]User),
		Transactions: make(map[int]Transaction),
	}
}

func (w *WalletService) AddUser(name string, email string) int {
	w.mu.Lock()
	defer w.mu.Unlock()

	id := len(w.Users) + 1
	w.Users[id] = User{
		id:    id,
		name:  name,
		email: email,
	}
	return id
}

func (w *WalletService) AddWallet(userId int, amount float64) int {
	w.mu.Lock()
	defer w.mu.Unlock()

	id := len(w.Wallets) + 1
	w.Wallets[userId] = &Wallet{
		id:     id,
		userId: userId,
		amount: amount,
	}
	return id
}

func (w *WalletService) TransferMoney(fromUserId int, toUserId int, amount float64) error {
	w.mu.RLock()
	if _, ok := w.Users[fromUserId]; !ok {
		return fmt.Errorf("No user %d", fromUserId)
	}

	if _, ok := w.Users[toUserId]; !ok {
		return fmt.Errorf("No user %d", toUserId)
	}

	fromUserWallet, ok := w.Wallets[fromUserId]
	if !ok {
		return fmt.Errorf("No wallet for user id %d", fromUserId)
	}

	toUserWallet, ok := w.Wallets[toUserId]
	if !ok {
		return fmt.Errorf("No wallet for user id %d", toUserId)
	}
	w.mu.RUnlock()

	first, second := fromUserWallet, toUserWallet
	if fromUserWallet.id > toUserWallet.id {
		first, second = toUserWallet, fromUserWallet
	}

	first.mu.Lock()
	defer first.mu.Unlock()
	second.mu.Lock()
	defer second.mu.Unlock()

	if fromUserWallet.amount < amount {
		return fmt.Errorf("Not enough money in wallet of user id %d", fromUserId)
	}

	fromUserWallet.amount -= amount
	toUserWallet.amount += amount

	w.mu.Lock()
	defer w.mu.Unlock()
	tid := len(w.Transactions) + 1
	w.Transactions[tid] = Transaction{
		Id:         tid,
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		Amount:     amount,
	}
	return nil
}

func (w *WalletService) GetUserWalletAmount(userId int) (float64, error) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if _, ok := w.Users[userId]; !ok {
		return 0, fmt.Errorf("No user %d", userId)
	}

	userWallet, ok := w.Wallets[userId]
	if !ok {
		return 0, fmt.Errorf("No wallet for user id %d", userId)
	}

	userWallet.mu.RLock()
	defer userWallet.mu.RUnlock()
	return userWallet.amount, nil
}
