package miyoushe

import (
	"net/url"
	"strings"

	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/miyoushe-task/config"
)

type BBSHome struct {
	Navigator []*BBSNavigator `json:"navigator"`
}

type BBSNavigator struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	AppPath string `json:"app_path"`
}

func (h *BBSHome) GetLunaActId() string {
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

func GetBBSHome(gameId string) (*BBSHome, error) {
	return Exec[*BBSHome](R().SetQueryParam("gids", gameId), "GET", AddrBBS+"/apihub/api/home/new")
}

type SignBBSData struct {
	Points int `json:"points"`
}

func SignBBS(gameId string, account config.Account, validate *Validate) (*SignBBSData, error) {
	return Exec[*SignBBSData](R(account.Device, validate).SetCookies(hcSToken(account)).SetBody(gh.M{"gids": gameId}), "POST", AddrBBS+"/apihub/app/api/signIn", 2)
}
