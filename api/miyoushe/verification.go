package miyoushe

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

type CreateVerificationData struct {
	Challenge  string `json:"challenge"`
	Gt         string `json:"gt"`
	NewCaptcha int    `json:"new_captcha"`
	Success    int    `json:"success"`
}

func CreateVerification(account config.Account) (*CreateVerificationData, error) {
	req := common.R(account.Device).SetCookies(common.SToken(account))
	return common.Exec[*CreateVerificationData](req, "GET", AddrBBS+"/misc/api/createVerification?is_high=true")
}

type VerifyVerificationData struct {
	Challenge string `json:"challenge"`
}

func VerifyVerification(challenge, validate string, account config.Account) (*VerifyVerificationData, error) {
	body := gh.M{"geetest_challenge": challenge, "geetest_seccode": validate + "|jordan", "geetest_validate": validate}
	req := common.R(account.Device).SetCookies(common.SToken(account)).SetBody(body)
	return common.Exec[*VerifyVerificationData](req, "POST", AddrBBS+"/misc/api/verifyVerification")
}
