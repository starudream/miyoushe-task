package common

import (
	"net/http"

	"github.com/starudream/miyoushe-task/config"
)

// SToken generate base stoken auth cookies
func SToken(account config.Account) []*http.Cookie {
	return []*http.Cookie{
		{Name: "mid", Value: account.Mid},
		{Name: "stoken", Value: account.SToken},
	}
}

// CToken generate base cookie_token auth cookies
func CToken(account config.Account) []*http.Cookie {
	return []*http.Cookie{
		{Name: "account_id", Value: account.Uid},
		{Name: "cookie_token", Value: account.CToken},
	}
}
