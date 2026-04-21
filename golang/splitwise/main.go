package main

import (
	"fmt"
	"splitwise/internal"
)

func main() {
	splitwiseService := internal.NewSplitService()
	splitwiseService.AddUser(1, "Alice", "1234567890", "alice@example.com")
	splitwiseService.AddUser(2, "Bob", "0987654321", "bob@example.com")

	groupId := splitwiseService.AddGroup("Trip to Paris", []int{1, 2})

	splitwiseService.AddExpense(groupId, "Hotel", 200.0, 1, "equal", []internal.Split{
		{UserId: 1},
		{UserId: 2},
	})

	splitwiseService.AddExpense(groupId, "Dinner", 100.0, 2, "equal", []internal.Split{
		{UserId: 1},
		{UserId: 2},
	})

	userId := 1
	balances := splitwiseService.GetBalance(userId)
	for user, amount := range balances {
		if userId != user {
			fmt.Printf("User %d owes User %d: %.2f\n", userId, user, amount)
		}
	}
}
