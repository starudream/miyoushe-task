package mihoyo

import (
	"fmt"
	"net/url"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

type GenQRCodeData struct {
	Url    string `json:"url"`
	Ticket string `json:"ticket"`
}

// GenQRCode generate to login sr (app_id = 8) to get game token
// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/hoyolab/login/qrcode_hk4e.md#%E7%94%9F%E6%88%90%E4%BA%8C%E7%BB%B4%E7%A0%81
func GenQRCode(account config.Account) (*GenQRCodeData, error) {
	req := common.R(account.Device).SetBody(gh.M{"app_id": "8", "device": account.Device.Id})
	data, err := common.Exec[*GenQRCodeData](req, "POST", AddrHK4E+"/hk4e_cn/combo/panda/qrcode/fetch")
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
	req := common.R(account.Device).SetBody(gh.M{"app_id": "8", "device": account.Device.Id, "ticket": ticket})
	data, err := common.Exec[*QueryQRCodeData](req, "POST", AddrHK4E+"/hk4e_cn/combo/panda/qrcode/query")
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
