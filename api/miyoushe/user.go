package miyoushe

import (
	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

func Login(account config.Account) error {
	req := common.R(account.Device).SetCookies(common.SToken(account)).SetBody(gh.M{"source_id": "", "source_key": "", "source_name": "", "source_type": 0})
	_, err := common.Exec[any](req, "POST", AddrBBS+"/user/api/login")
	return err
}

type GetUserData struct {
	UserInfo *UserInfo `json:"user_info"`
}

type UserInfo struct {
	Uid       string `json:"uid"`
	Nickname  string `json:"nickname"`
	Introduce string `json:"introduce"`
}

func GetUser(uid string, account ...config.Account) (*GetUserData, error) {
	var req *resty.Request
	if len(account) == 0 {
		req = common.R().SetQueryParam("uid", uid)
	} else {
		req = common.R(account[0].Device).SetCookies(common.SToken(account[0]))
	}
	return common.Exec[*GetUserData](req, "GET", AddrBBS+"/user/api/getUserFullInfo")
}
