package test

import (
	"digital-wallet/internal"
	"log"
	"sync"
	"testing"
)

func TestUserTransaction(t *testing.T) {
	walletService := internal.NewWalletService()
	userA := walletService.AddUser("vd", "vd@gmail.com")
	userB := walletService.AddUser("rj", "rj@gmail.com")
	walletService.AddWallet(userA, 2900)
	walletService.AddWallet(userB, 1000)
	err := walletService.TransferMoney(userA, userB, 1000)
	if err != nil {
		t.Fatal(err)
	}

	amount, err := walletService.GetUserWalletAmount(userA)
	if err != nil {
		t.Fatal(err)
	}
	if amount != 1900 {
		t.Errorf("Invalid amount for user %v", amount)
	}

	amount, err = walletService.GetUserWalletAmount(userB)
	if err != nil {
		log.Printf("Error : %v", err)
	}
	if amount != 2000 {
		t.Errorf("Invalid amount for user %v", amount)
	}
}

func TestConcurrent(t *testing.T) {
	walletService := internal.NewWalletService()
	userA := walletService.AddUser("vd", "vd@gmail.com")
	userB := walletService.AddUser("rj", "rj@gmail.com")
	walletService.AddWallet(userA, 2900)
	walletService.AddWallet(userB, 1000)

	const workers = 10
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				walletService.TransferMoney(userA, userB, 1000)
			} else {
				walletService.TransferMoney(userB, userA, 1000)
			}
		}(i)
	}

	wg.Wait()
	amount, err := walletService.GetUserWalletAmount(userA)
	if err != nil {
		t.Fatal(err)
	}
	if amount != 2900 {
		t.Errorf("Invalid amount for user %v", amount)
	}

	amount, err = walletService.GetUserWalletAmount(userB)
	if err != nil {
		log.Printf("Error : %v", err)
	}
	if amount != 1000 {
		t.Errorf("Invalid amount for user %v", amount)
	}
}

func TestConcurrencyAndDeadlock(t *testing.T) {
	service := internal.NewWalletService()

	user1 := service.AddUser("Alice", "alice@test.com")
	user2 := service.AddUser("Bob", "bob@test.com")

	service.AddWallet(user1, 10000)
	service.AddWallet(user2, 10000)

	const iterations = 1000
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			_ = service.TransferMoney(user1, user2, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			_ = service.TransferMoney(user2, user1, 1)
		}
	}()

	wg.Wait()

	bal1, _ := service.GetUserWalletAmount(user1)
	bal2, _ := service.GetUserWalletAmount(user2)

	if bal1 != 10000 || bal2 != 10000 {
		t.Errorf("Consistency error: Alice: %v, Bob: %v", bal1, bal2)
	}
}
