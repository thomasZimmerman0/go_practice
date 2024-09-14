package main

import (
	"fmt"
	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main(){
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		// panic("Can't continue, sorry.")
	}

	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice:")
		fmt.Scan(&choice)
	
		if choice == 1 {
				fmt.Println("Your balance is", accountBalance)
		} else if choice ==2 {
			fmt.Print("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//return
				continue
			} 
	
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated!New amount:", accountBalance)
		} else if choice == 3 {
			fmt.Print("Withdrawl amount: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)
	
			if withdrawlAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//return
				continue
			} 
	
			if withdrawlAmount > accountBalance {
				fmt.Println("You cannot withdrawl more than you have.")
				//return
				continue
			} 
	
			accountBalance -= withdrawlAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			//return 
			break;
		}
	
		fmt.Println("Your choice:", choice)
	}

	fmt.Println("Thanks for chosing our bank")
}
