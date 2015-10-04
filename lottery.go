package main

import (
	"github.com/deet/govenmo"
	"github.com/jmcvetta/randutil"
	"log"
)

func main() {
	account, err := venmoAccount()
	if err != nil {
		log.Println("Error refreshing account:", err)
	} else {
		log.Println("Account refreshed")
	}

	log.Println("balance is", account.Balance)
	payments, err := paymentsSinceLastRun(&account)
	if err != nil {
		log.Println("Error fetching payments:", err)
		return
	}

	balance := 0.0

	choices := make([]randutil.Choice, 0)

	for _, payment := range payments {
		actor := payment.Actor.DisplayName
		target := payment.Target.User.DisplayName
		note := payment.Note
		amount := payment.Amount
		log.Println(actor, "paid", target, amount, "dollars for", note)

		if target == "Smick Share" {
			balance += amount
			choice := randutil.Choice{
				Weight: int(100 * amount),
				Item:   payment.Actor.Id,
			}
			choices = append(choices, choice)
		}
	}

	if balance == 0 {
		log.Println("no money to give out")
		return
	}

	log.Println("Final balance is", balance)
	log.Println(choices)
	winner, err := randutil.WeightedChoice(choices)
	if err != nil {
		log.Println("Error selecting winner:", err)
		return
	}
	winnerId := winner.Item

	target := govenmo.Target{}
	target.User.Id = winnerId.(string)

	sentPayment, err := account.PayOrCharge(target, balance, "ayy lmao!", "public")
	if err != nil {
		log.Println("Error sending payment:", err)
	} else {
		log.Println("Payment succeeded", sentPayment)
	}
}
