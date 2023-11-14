package miyoushe

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

type SignForumData struct {
	Points int `json:"points"`
}

func SignForum(gameId string, account config.Account, verification *common.Verification) (*SignForumData, error) {
	req := common.R(account.Device, verification).SetCookies(common.SToken(account)).SetBody(gh.M{"gids": gameId})
	return common.Exec[*SignForumData](req, "POST", AddrBBS+"/apihub/app/api/signIn", 2)
}

type GetSignForumData struct {
	IsSigned bool `json:"is_signed"`
}

func GetSignForum(gameId string, account config.Account) (*GetSignForumData, error) {
	req := common.R(account.Device).SetCookies(common.SToken(account)).SetQueryParam("gids", gameId)
	return common.Exec[*GetSignForumData](req, "GET", AddrBBS+"/apihub/sapi/querySignInStatus")
}
