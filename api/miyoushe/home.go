package miyoushe

import (
	"net/url"
	"strings"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

type Home struct {
	Navigator []*HomeNav    `json:"navigator"`
	Official  *HomeOfficial `json:"official"`
}

type HomeNav struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AppPath string `json:"app_path"`
}

type HomeOfficial struct {
	ForumId int `json:"forum_id"`
}

func (h *Home) GetSignActId() string {
	for _, nav := range h.Navigator {
		if strings.Contains(nav.Name, "签到") {
			u, err := url.Parse(nav.AppPath)
			if err == nil {
				return u.Query().Get("act_id")
			}
		}
	}
	return ""
}

func GetHome(gameId string, account config.Account) (*Home, error) {
	req := common.R(account.Device).SetQueryParam("gids", gameId)
	return common.Exec[*Home](req, "GET", AddrBBS+"/apihub/api/home/new")
}

type GetBusinessesData struct {
	Businesses []string `json:"businesses"`
}

func GetBusinesses(account config.Account) (*GetBusinessesData, error) {
	req := common.R(account.Device).SetCookies(common.SToken(account)).SetQueryParam("uid", account.Uid)
	return common.Exec[*GetBusinessesData](req, "GET", AddrBBS+"/user/api/getUserBusinesses")
}
