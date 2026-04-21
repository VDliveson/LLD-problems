package tests

import (
	"fmt"
	"splitwise/internal"
	"sync"
	"testing"
)

func TestSplitwise(t *testing.T) {
	splitwiseService := internal.NewSplitService()
	splitwiseService.AddUser(1, "Alice", "1234567890", "alice@example.com")
	splitwiseService.AddUser(2, "Bob", "0987654321", "bob@example.com")
	splitwiseService.AddGroup("Trip to Paris", []int{1, 2})

	splitwiseService.AddExpense(1, "Hotel", 200.0, 1, "equal", []internal.Split{
		{UserId: 1},
		{UserId: 2},
	})
	splitwiseService.AddExpense(1, "Dinner", 100.0, 2, "equal", []internal.Split{
		{UserId: 1},
		{UserId: 2},
	})

	balances := splitwiseService.GetBalance(1)
	if balances[2] != -50.0 {
		t.Errorf("Expected User 1 to owe User 2: -50.0, got: %.2f", balances[2])
	}

	balances = splitwiseService.GetBalance(2)
	if balances[1] != 50.0 {
		t.Errorf("Expected User 2 to owe User 1: 50.0, got: %.2f", balances[1])
	}
}

func TestSplitwiseConcurrent(t *testing.T) {
	splitwiseService := internal.NewSplitService()

	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(userId int) {
			defer wg.Done()
			splitwiseService.AddUser(userId, fmt.Sprintf("User%d", userId), fmt.Sprintf("123456789%d", userId), fmt.Sprintf("user%d@example.com", userId))
		}(i)
	}
	wg.Wait()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(groupId int) {
			defer wg.Done()
			splitwiseService.AddGroup(fmt.Sprintf("Group%d", groupId), []int{i, i + 1, i + 2})
		}(i)
	}
	wg.Wait()

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			splits := []internal.Split{
				{UserId: i},
				{UserId: i + 1},
			}
			splitwiseService.AddExpense(i, fmt.Sprintf("Expense%d", i), float64(i*10), i, "equal", splits)
		}(i)
	}
	wg.Wait()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(userId int) {
			defer wg.Done()
			balances := splitwiseService.GetBalance(userId)
			fmt.Printf("User %d balances: %v\n", userId, balances)
		}(i)
	}
}
