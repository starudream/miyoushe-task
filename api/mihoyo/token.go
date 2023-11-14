package mihoyo

import (
	"fmt"
	"strconv"

	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

type GetSTokenByGTokenData struct {
	Token    *TokenInfo `json:"token"`
	UserInfo *UserInfo  `json:"user_info"`
}

type TokenInfo struct {
	TokenType int    `json:"token_type"`
	Token     string `json:"token"`
}

type UserInfo struct {
	Aid string `json:"aid"`
	Mid string `json:"mid"`
}

// GetSTokenByGToken get stoken v2 by game token
// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/hoyolab/user/token.md#%E9%80%9A%E8%BF%87game-token%E8%8E%B7%E5%8F%96stokenv1
func GetSTokenByGToken(account config.Account) (*GetSTokenByGTokenData, error) {
	uid, err := strconv.ParseInt(account.Uid, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse uid error: %w", err)
	}
	req := common.R(account.Device).SetBody(gh.M{"account_id": uid, "game_token": account.GToken})
	return common.Exec[*GetSTokenByGTokenData](req, "POST", AddrTakumi+"/account/ma-cn-session/app/getTokenByGameToken")
}

type GetCTokenBySTokenData struct {
	Uid         string `json:"uid"`
	CookieToken string `json:"cookie_token"`
}

// GetCTokenBySToken get cookie token by stoken v2
// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/hoyolab/user/token.md#%E9%80%9A%E8%BF%87stoken%E8%8E%B7%E5%8F%96cookie-token
func GetCTokenBySToken(account config.Account) (*GetCTokenBySTokenData, error) {
	req := common.R(account.Device).SetCookies(common.SToken(account))
	return common.Exec[*GetCTokenBySTokenData](req, "GET", AddrTakumi+"/auth/api/getCookieAccountInfoBySToken")
}

type GetLTokenBySTokenData struct {
	LToken string `json:"ltoken"`
}

// GetLTokenBySToken get ltoken v1 by stoken v2
// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/hoyolab/user/token.md#%E9%80%9A%E8%BF%87stoken%E8%8E%B7%E5%8F%96ltokenv1
func GetLTokenBySToken(account config.Account) (*GetLTokenBySTokenData, error) {
	req := common.R(account.Device).SetCookies(common.SToken(account))
	return common.Exec[*GetLTokenBySTokenData](req, "GET", AddrTakumi+"/account/auth/api/getLTokenBySToken")
}
