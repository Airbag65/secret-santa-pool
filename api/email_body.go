package main

import "fmt"


func SwedishLotteryMail(buyer, reciever *Person, pool *Pool) (string, []string) {
	return fmt.Sprintf("Lottningen är gjord %s %s!!", buyer.FirstName, buyer.FirstName), []string{
		fmt.Sprintf("Hej %s %s!\n\n"),
		fmt.Sprintf("Lottningen är nu utförd! Du har dragit... %s %s att köpa en present till.\r\n", reciever.FirstName, reciever.LastName),
		fmt.Sprintf("Observera att den bestämda summan att handla för är %d %s i denna pool\r\n", pool.Amount, pool.Currency),
		fmt.Sprintf("Lycka till, och god jul!"),
	}
}

func EnglishLotteryMail(buyer, reciever *Person, pool *Pool) (string, []string) {
	return fmt.Sprintf("The lottery is done %s %s!!", buyer.FirstName, buyer.FirstName), []string{
		fmt.Sprintf("Hello %s %s!\n\n"),
		fmt.Sprintf("The lottery has now been performed! You have drawn... %s %s to buy for.\r\n", reciever.FirstName, reciever.LastName),
		fmt.Sprintf("Keep in mind that the agreed amount to buy for in this pool is %d %s\r\n", pool.Amount, pool.Currency),
		fmt.Sprintf("Good luck, and happy holidays!"),
	}
}

