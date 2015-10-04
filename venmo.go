package main

import (
	"fmt"
	"github.com/deet/govenmo"
	"github.com/jmcvetta/randutil"
	"time"
)

func main() {
	account := govenmo.Account{
		AccessToken: "c25c552dba68c004af633f6535fe57c8f8c8c06c9eca5b140aa8974eb3067ff9",
	}

	if err := account.Refresh(); err != nil {
		fmt.Println("Error refreshing account:", err)
	} else {
		fmt.Println("Account refreshed")
	}

	fmt.Println("balance is", account.Balance)

	updatedSince := time.Now()
	payments, err := account.PaymentsSince(updatedSince)
	if err != nil {
		fmt.Println("Error fetching payments:", err)
		return
	}

	balance := 0.0

	choices := make([]randutil.Choice, 0)

	for _, payment := range payments {
		actor := payment.Actor.DisplayName
		target := payment.Target.User.DisplayName
		note := payment.Note
		amount := payment.Amount
		fmt.Println(actor, "paid", target, amount, "dollars for", note)

		if target == "Smick Share" {
			balance += amount
			choice := randutil.Choice{
				Weight: int(100 * amount),
				Item:   payment.Actor.Id,
			}
			choices = append(choices, choice)
		} else {
			break
		}
	}

	if balance == 0 {
		fmt.Println("no money to give out")
		return
	}

	fmt.Println("Final balance is", balance)
	fmt.Println(choices)
	winner, err := randutil.WeightedChoice(choices)
	if err != nil {
		fmt.Println("Error selecting winner:", err)
		return
	}
	winnerId := winner.Item

	target := govenmo.Target{}
	target.User.Id = winnerId.(string)

	sentPayment, err := account.PayOrCharge(target, balance, "ayy lmao!", "public")
	if err != nil {
		fmt.Println("Error sending payment:", err)
	} else {
		fmt.Println("Payment succeeded", sentPayment)
	}
}
