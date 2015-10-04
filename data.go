package main

import (
	"github.com/deet/govenmo"
	"time"
)

const VenmoDisplayName = "Smick Share"

func paymentsSinceLastRun(account *govenmo.Account) (payments []govenmo.Payment, err error) {
	payments, err = account.PaymentsSince(time.Now())
	if err != nil {
		return
	}

	for i, payment := range payments {
		target := payment.Target.User.DisplayName
		if target != VenmoDisplayName {
			return payments[:i], err
		}
	}

	return
}
