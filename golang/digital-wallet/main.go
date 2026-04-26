package main

import (
	"digital-wallet/internal"
	"log"
)

func main() {
	walletService := internal.NewWalletService()
	userA := walletService.AddUser("vd", "vd@gmail.com")
	userB := walletService.AddUser("rj", "rj@gmail.com")
	walletService.AddWallet(userA, 2900)
	walletService.AddWallet(userB, 1000)
	err := walletService.TransferMoney(userA, userB, 1000)
	if err != nil {
		log.Printf("Error : %v", err)
	}

	amount, err := walletService.GetUserWalletAmount(userA)
	if err != nil {
		log.Printf("Error : %v", err)
	}
	log.Printf("Amount for user id: %d is %v\n", userA, amount)

	amount, err = walletService.GetUserWalletAmount(userB)
	if err != nil {
		log.Printf("Error : %v", err)
	}
	log.Printf("Amount for user id: %d is %v\n", userB, amount)
}
