package miyoushe

import (
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/miyoushe-task/config"
)

type User struct {
	UserInfo *UserInfo `json:"user_info"`
}

type UserInfo struct {
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`

	Aid           string `json:"aid"`
	Mid           string `json:"mid"`
	AccountName   string `json:"account_name"`
	Email         string `json:"email"`
	IsEmailVerify int    `json:"is_email_verify"`
	AreaCode      string `json:"area_code"`
	Mobile        string `json:"mobile"`
	Realname      string `json:"realname"`
	IdentityCode  string `json:"identity_code"`
}

func GetUser(uid string, account ...config.Account) (*User, error) {
	var req *resty.Request
	if len(account) == 0 {
		req = R().SetQueryParam("uid", uid)
	} else {
		req = R(account[0].Device).SetCookies(hcSToken(account[0]))
	}
	return Exec[*User](req, "GET", AddrBBS+"/user/api/getUserFullInfo")
}
