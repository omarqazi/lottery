package main

import (
	"github.com/deet/govenmo"
)

func venmoAccount() (account govenmo.Account, err error) {
	account = govenmo.Account{
		AccessToken: venmoAccessToken(),
	}

	err = account.Refresh()
	return
}

func venmoAccessToken() string {
	return "393a7717ce46b7ad854a19473fa0b001313f2d2c54ddcca0f54b3c7657c80b07"
}
