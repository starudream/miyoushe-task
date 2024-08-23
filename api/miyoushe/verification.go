package miyoushe

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

func CreateVerification(account config.Account) (*common.AigisData, error) {
	req := common.R(account.Device).SetCookies(common.SToken(account)).SetQueryParam("is_high", "true")
	return common.Exec[*common.AigisData](req, "GET", AddrBBS+"/misc/api/createVerification")
}

type VerifyVerificationData struct {
	Challenge string `json:"challenge"`
}

func VerifyVerification(challenge, validate string, account config.Account) (*VerifyVerificationData, error) {
	body := gh.M{"geetest_challenge": challenge, "geetest_seccode": validate + "|jordan", "geetest_validate": validate}
	req := common.R(account.Device).SetCookies(common.SToken(account)).SetBody(body)
	return common.Exec[*VerifyVerificationData](req, "POST", AddrBBS+"/misc/api/verifyVerification")
}
