package miyoushe

import (
	"cmp"
	"slices"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

type ListGameData struct {
	List []*Game `json:"list"`
}

type Game struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	EnName string `json:"en_name"`
	OpName string `json:"op_name"`
}

func ListGame() (*ListGameData, error) {
	data, err := common.Exec[*ListGameData](common.R(), "GET", AddrBBS+"/apihub/api/getGameList")
	if err != nil {
		return nil, err
	}
	slices.SortFunc(data.List, func(a, b *Game) int { return cmp.Compare(a.Id, b.Id) })
	return data, nil
}

type ListGameRoleData struct {
	List []*GameRole `json:"list"`
}

type GameRole struct {
	GameBiz    string `json:"game_biz"`
	Region     string `json:"region"`
	GameUid    string `json:"game_uid"`
	Nickname   string `json:"nickname"`
	Level      int    `json:"level"`
	IsChosen   bool   `json:"is_chosen"`
	RegionName string `json:"region_name"`
	IsOfficial bool   `json:"is_official"`
}

func ListGameRole(gameBiz string, account config.Account) (*ListGameRoleData, error) {
	req := common.R(account.Device).SetCookies(common.SToken(account)).SetQueryParam("game_biz", gameBiz)
	return common.Exec[*ListGameRoleData](req, "GET", AddrTakumi+"/binding/api/getUserGameRolesByStoken")
}
