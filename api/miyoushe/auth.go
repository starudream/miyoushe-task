package miyoushe

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/miyoushe-task/config"
)

type GenQRCodeData struct {
	Url    string `json:"url"`
	Ticket string `json:"ticket"`
}

// GenQRCode generate to login sr (app_id = 8) to get game token
// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/hoyolab/login/qrcode_hk4e.md#%E7%94%9F%E6%88%90%E4%BA%8C%E7%BB%B4%E7%A0%81
func GenQRCode(account config.Account) (*GenQRCodeData, error) {
	data, err := Exec[*GenQRCodeData](R(account.Device).SetBody(gh.M{"app_id": "8", "device": account.Device.Id}), "POST", AddrHK4E+"/hk4e_cn/combo/panda/qrcode/fetch")
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(data.Url)
	if err != nil {
		return nil, fmt.Errorf("parse url error: %w", err)
	}
	data.Ticket = u.Query().Get("ticket")
	return data, nil
}

type QueryQRCodeData struct {
	Stat    QRCodeStat          `json:"stat"`
	Payload *QueryQRCodePayload `json:"payload"`
}

type QueryQRCodePayload struct {
	Proto string `json:"proto"`
	Raw   string `json:"raw,omitempty"`
	Uid   string `json:"uid,omitempty"`
	Token string `json:"token,omitempty"`
}

type QRCodeStat string

const (
	QRCodeStatInit      = "Init"
	QRCodeStatScanned   = "Scanned"
	QRCodeStatConfirmed = "Confirmed"
)

func (s QRCodeStat) IsInit() bool {
	return s == QRCodeStatInit
}

func (s QRCodeStat) IsScanned() bool {
	return s == QRCodeStatScanned
}

func (s QRCodeStat) IsConfirmed() bool {
	return s == QRCodeStatConfirmed
}

func QueryQRCode(ticket string, account config.Account) (*QueryQRCodeData, error) {
	data, err := Exec[*QueryQRCodeData](R(account.Device).SetBody(gh.M{"app_id": "8", "device": account.Device.Id, "ticket": ticket}), "POST", AddrHK4E+"/hk4e_cn/combo/panda/qrcode/query")
	if err != nil || !data.Stat.IsConfirmed() {
		return data, err
	}
	payload, err := json.UnmarshalTo[*QueryQRCodePayload](data.Payload.Raw)
	if err != nil {
		return nil, fmt.Errorf("unmarshal payload error: %w", err)
	}
	data.Payload.Uid = payload.Uid
	data.Payload.Token = payload.Token // game token
	return data, nil
}

type GetSTokenByGTokenData struct {
	Token    *TokenInfo `json:"token"`
	UserInfo *UserInfo  `json:"user_info"`
}

type TokenInfo struct {
	TokenType int    `json:"token_type"`
	Token     string `json:"token"`
}

// GetSTokenByGToken get stoken v2 by game token
// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/hoyolab/user/token.md#%E9%80%9A%E8%BF%87game-token%E8%8E%B7%E5%8F%96stokenv1
func GetSTokenByGToken(account config.Account) (*GetSTokenByGTokenData, error) {
	uid, err := strconv.ParseInt(account.Uid, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse uid error: %w", err)
	}
	return Exec[*GetSTokenByGTokenData](R(account.Device).SetBody(gh.M{"account_id": uid, "game_token": account.GToken}), "POST", AddrTakumi+"/account/ma-cn-session/app/getTokenByGameToken")
}

type GetCTokenBySTokenData struct {
	Uid         string `json:"uid"`
	CookieToken string `json:"cookie_token"`
}

// GetCTokenBySToken get cookie token by stoken v2
// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/hoyolab/user/token.md#%E9%80%9A%E8%BF%87stoken%E8%8E%B7%E5%8F%96cookie-token
func GetCTokenBySToken(account config.Account) (*GetCTokenBySTokenData, error) {
	return Exec[*GetCTokenBySTokenData](R(account.Device).SetCookies(hcSToken(account)), "GET", AddrTakumi+"/auth/api/getCookieAccountInfoBySToken")
}

type GetLTokenBySTokenData struct {
	LToken string `json:"ltoken"`
}

// GetLTokenBySToken get ltoken v1 by stoken v2
// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/hoyolab/user/token.md#%E9%80%9A%E8%BF%87stoken%E8%8E%B7%E5%8F%96ltokenv1
func GetLTokenBySToken(account config.Account) (*GetLTokenBySTokenData, error) {
	return Exec[*GetLTokenBySTokenData](R(account.Device).SetCookies(hcSToken(account)), "GET", AddrTakumi+"/account/auth/api/getLTokenBySToken")
}
