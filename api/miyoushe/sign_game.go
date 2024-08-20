package miyoushe

import (
	"strconv"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

var signGameAddrByName = map[string]string{
	common.GameNameZZZ: AddrActNap,
}

func signGameAddr(gameName string) string {
	addr, ok := signGameAddrByName[gameName]
	if ok {
		return addr
	}
	return AddrTakumi
}

var signGameHeaderByName = map[string]string{
	common.GameNameZZZ: "zzz",
}

func signGameHeader(gameName string) string {
	header, ok := signGameHeaderByName[gameName]
	if ok {
		return header
	}
	return gameName
}

type SignGameData struct {
	Code      string `json:"code"`
	Success   int    `json:"success"`
	IsRisk    bool   `json:"is_risk"`
	RiskCode  int    `json:"risk_code"`
	Gt        string `json:"gt"`
	Challenge string `json:"challenge"`
}

func (t *SignGameData) IsRisky() bool {
	return t.IsRisk
}

func SignGame(gameName, actId, region, uid string, account config.Account, validate *common.Verification) (*SignGameData, error) {
	body := gh.MS{"lang": "zh-cn", "act_id": actId, "region": region, "uid": uid}
	req := common.R(account.Device, validate).SetHeader(common.XRpcSignGame, signGameHeader(gameName)).SetCookies(common.SToken(account)).SetCookies(common.CToken(account)).SetBody(body)
	return common.Exec[*SignGameData](req, "POST", signGameAddr(gameName)+"/event/luna/sign")
}

type GetSignGameData struct {
	TotalSignDay  int    `json:"total_sign_day"`
	Today         string `json:"today"`
	IsSign        bool   `json:"is_sign"`
	IsSub         bool   `json:"is_sub"`
	Region        string `json:"region"`
	SignCntMissed int    `json:"sign_cnt_missed"`
	ShortSignDay  int    `json:"short_sign_day"`
}

func GetSignGame(gameName, actId, region, uid string, account config.Account) (*GetSignGameData, error) {
	query := gh.MS{"lang": "zh-cn", "act_id": actId, "region": region, "uid": uid}
	req := common.R(account.Device).SetHeader(common.XRpcSignGame, signGameHeader(gameName)).SetCookies(common.SToken(account)).SetCookies(common.CToken(account)).SetQueryParams(query)
	return common.Exec[*GetSignGameData](req, "GET", signGameAddr(gameName)+"/event/luna/info")
}

type ListSignGameData struct {
	Month      int                 `json:"month"`
	Biz        string              `json:"biz"`
	Resign     bool                `json:"resign"`
	Awards     []*SignGameAward    `json:"awards"`
	ExtraAward *SignGameExtraAward `json:"short_extra_award"`
}

type SignGameAward struct {
	Name      string `json:"name"`
	Cnt       int    `json:"cnt"`
	CreatedAt string `json:"created_at,omitempty"`
}

type SignGameExtraAward struct {
	HasExtraAward  bool   `json:"has_extra_award"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	List           []any  `json:"list"`
	StartTimestamp string `json:"start_timestamp"`
	EndTimestamp   string `json:"end_timestamp"`
}

func ListSignGame(gameName, actId string, account config.Account) (*ListSignGameData, error) {
	query := gh.MS{"lang": "zh-cn", "act_id": actId}
	req := common.R(account.Device).SetHeader(common.XRpcSignGame, signGameHeader(gameName)).SetCookies(common.SToken(account)).SetQueryParams(query)
	return common.Exec[*ListSignGameData](req, "GET", signGameAddr(gameName)+"/event/luna/home")
}

type ListSignGameAwardData struct {
	Total int            `json:"total"`
	List  SignGameAwards `json:"list"`
}

type SignGameAwards []*SignGameAward

func (v1 SignGameAwards) Today() (v2 SignGameAwards) {
	today := time.Now().Format(time.DateOnly)
	for i := range v1 {
		if strings.HasPrefix(v1[i].CreatedAt, today) {
			v2 = append(v2, v1[i])
		}
	}
	return
}

func (v1 SignGameAwards) ShortString() string {
	v2 := make([]string, len(v1))
	for i, v := range v1 {
		v2[i] = v.Name + "*" + strconv.Itoa(v.Cnt)
	}
	return strings.Join(v2, ", ")
}

func ListSignGameAward(gameName, actId, region, uid string, account config.Account) (list SignGameAwards, _ error) {
	for page, total := 1, -1; ; page++ {
		query := gh.MS{"lang": "zh-cn", "act_id": actId, "region": region, "uid": uid, "current_page": strconv.Itoa(page), "page_size": "10"}
		req := common.R(account.Device).SetHeader(common.XRpcSignGame, signGameHeader(gameName)).SetCookies(common.SToken(account)).SetCookies(common.CToken(account)).SetQueryParams(query)
		data, err := common.Exec[*ListSignGameAwardData](req, "GET", signGameAddr(gameName)+"/event/luna/award", 2)
		if err != nil {
			return nil, err
		}
		list = append(list, data.List...)
		if page == 1 && total == -1 {
			total = data.Total
		}
		total -= len(data.List)
		if total <= 0 {
			return
		}
	}
}
