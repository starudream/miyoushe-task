package miyoushe

import (
	"strconv"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/miyoushe-task/config"
)

type SignLunaData struct {
	Code      string `json:"code"`
	Success   int    `json:"success"`
	IsRisk    bool   `json:"is_risk"`
	RiskCode  int    `json:"risk_code"`
	Challenge string `json:"challenge"`
	Gt        string `json:"gt"`
}

func (t *SignLunaData) IsSuccess() bool {
	return !t.IsRisk && t.RiskCode == 0
}

func SignLuna(actId, region, uid string, account config.Account) (*SignLunaData, error) {
	body := gh.MS{"lang": LangZHCN, "act_id": actId, "region": region, "uid": uid}
	return Exec[*SignLunaData](R(account.Device).SetCookies(hcSToken(account)).SetCookies(hcCToken(account)).SetBody(body), "POST", AddrTakumi+"/event/luna/sign")
}

type GetLunaTodayData struct {
	TotalSignDay  int    `json:"total_sign_day"`
	Today         string `json:"today"`
	IsSign        bool   `json:"is_sign"`
	IsSub         bool   `json:"is_sub"`
	Region        string `json:"region"`
	SignCntMissed int    `json:"sign_cnt_missed"`
	ShortSignDay  int    `json:"short_sign_day"`
}

func GetLunaToday(actId, region, uid string, account config.Account) (*GetLunaTodayData, error) {
	query := gh.MS{"lang": LangZHCN, "act_id": actId, "region": region, "uid": uid}
	return Exec[*GetLunaTodayData](R(account.Device).SetCookies(hcSToken(account)).SetCookies(hcCToken(account)).SetQueryParams(query), "GET", AddrTakumi+"/event/luna/info")
}

type ListLunaData struct {
	Month      int             `json:"month"`
	Biz        string          `json:"biz"`
	Resign     bool            `json:"resign"`
	Awards     []*LunaAward    `json:"awards"`
	ExtraAward *LunaExtraAward `json:"short_extra_award"`
}

type LunaAward struct {
	Name      string `json:"name"`
	Cnt       int    `json:"cnt"`
	CreatedAt string `json:"created_at,omitempty"`
}

type LunaExtraAward struct {
	HasExtraAward  bool   `json:"has_extra_award"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	List           []any  `json:"list"`
	StartTimestamp string `json:"start_timestamp"`
	EndTimestamp   string `json:"end_timestamp"`
}

func ListLuna(actId string, account config.Account) (*ListLunaData, error) {
	query := gh.MS{"lang": LangZHCN, "act_id": actId}
	return Exec[*ListLunaData](R(account.Device).SetCookies(hcSToken(account)).SetQueryParams(query), "GET", AddrTakumi+"/event/luna/home")
}

type ListLunaAwardData struct {
	Total int        `json:"total"`
	List  LunaAwards `json:"list"`
}

type LunaAwards []*LunaAward

func (v1 LunaAwards) Today() (v2 LunaAwards) {
	today := time.Now().Format(time.DateOnly)
	for i := range v1 {
		if strings.HasPrefix(v1[i].CreatedAt, today) {
			v2 = append(v2, v1[i])
		}
	}
	return
}

func (v1 LunaAwards) ShortString() string {
	v2 := make([]string, len(v1))
	for i, v := range v1 {
		v2[i] = v.Name + "*" + strconv.Itoa(v.Cnt)
	}
	return strings.Join(v2, ", ")
}

func ListLunaAward(actId, region, uid string, account config.Account) (list LunaAwards, _ error) {
	for page, total := 1, -1; ; page++ {
		query := gh.MS{"lang": LangZHCN, "act_id": actId, "region": region, "uid": uid, "current_page": strconv.Itoa(page), "page_size": "10"}
		data, err := Exec[*ListLunaAwardData](R(account.Device).SetCookies(hcSToken(account)).SetCookies(hcCToken(account)).SetQueryParams(query), "GET", AddrTakumi+"/event/luna/award")
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
