package miyoushe

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/miyoushe-task/config"
)

const (
	AddrBBS    = "https://bbs-api.miyoushe.com"
	AddrTakumi = "https://api-takumi.mihoyo.com"
	AddrHK4E   = "https://hk4e-sdk.mihoyo.com"

	UserAgent  = "Mozilla/5.0 (Linux; Android 13; 22011211C Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/104.0.5112.97 Mobile Safari/537.36 miHoYoBBS/2.62.2"
	Referer    = "https://app.mihoyo.com"
	RefererAct = "https://act.mihoyo.com"

	AppVersion = "2.62.2"
	// AppSaltK2 https://blog.starudream.cn/2023/11/09/miyoushe-salt-2.62.2/
	AppSaltK2 = "pIlzNr5SAZhdnFW8ZxauW8UlxRdZc45r"
	// AppSalt6X https://github.com/UIGF-org/mihoyo-api-collect/issues/1
	AppSalt6X     = "t0qEgfub6cvueAPgR5m9aQWWVciEer7v"
	AppIdMiyoushe = "bll8iq97cem8" // 米游社

	LangZHCN = "zh-cn"

	GameBizHK4ECN  = "hk4e_cn"  // 原神	国服
	GameBizHKRPGCN = "hkrpg_cn" // 崩坏：星穹铁道	国服

	GameIdYS = "2" // 原神
	GameIdSR = "6" // 崩坏：星穹铁道

	RetCodeQRCodeExpired   = -106  // 二维码过期
	RetCodeLunaHasSigned   = -5003 // 已签到
	RetCodeBBSHasSigned    = 1008  // 打卡失败或重复打卡
	RetCodeBBSNeedValidate = 1034  // 需要验证码

	gameJSON = `[{"id":1,"name":"崩坏3","en_name":"bh3","op_name":"bh3"},{"id":2,"name":"原神","en_name":"ys","op_name":"hk4e"},{"id":3,"name":"崩坏学园2","en_name":"bh2","op_name":"bh2"},{"id":4,"name":"未定事件簿","en_name":"wd","op_name":"nxx"},{"id":5,"name":"大别野","en_name":"dby","op_name":"plat"},{"id":6,"name":"崩坏：星穹铁道","en_name":"sr","op_name":"hkrpg"},{"id":8,"name":"绝区零","en_name":"zzz","op_name":"nap"}]`
)

var (
	GameIdByBiz = map[string]string{
		// GameBizHK4ECN:  GameIdYS,
		GameBizHKRPGCN: GameIdSR,
	}
	GameList = json.MustUnmarshalTo[[]*Game](gameJSON)
	GameById = func() map[string]*Game {
		m := map[string]*Game{}
		for i := range GameList {
			m[strconv.Itoa(GameList[i].Id)] = GameList[i]
		}
		return m
	}()
)

type BaseResp[T any] struct {
	RetCode *int   `json:"retcode"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func (t *BaseResp[T]) GetRetCode() int {
	if t == nil || t.RetCode == nil {
		return 999999
	}
	return *t.RetCode
}

func (t *BaseResp[T]) IsSuccess() bool {
	return t != nil && t.RetCode != nil && *t.RetCode == 0
}

func (t *BaseResp[T]) String() string {
	if t == nil || t.RetCode == nil {
		return "<nil>"
	}
	return fmt.Sprintf("retcode: %d, message: %s, data: %v", *t.RetCode, t.Message, t.Data)
}

func IsRetCode(err error, rc int) bool {
	if err == nil {
		return false
	}
	e, ok1 := resty.AsRespErr(err)
	if ok1 {
		t, ok2 := e.Result().(interface{ GetRetCode() int })
		if ok2 {
			return t.GetRetCode() == rc
		}
	}
	return false
}

type Validate struct {
	Challenge string `json:"challenge"`
	Validate  string `json:"validate"`
}

func R(vs ...any) *resty.Request {
	r := resty.R().
		SetHeader("Accept-Encoding", "gzip").
		SetHeader("User-Agent", UserAgent).
		SetHeader("Referer", Referer).
		SetHeader("x-rpc-app_version", AppVersion).
		SetHeader("x-rpc-app_id", AppIdMiyoushe).
		SetHeader("x-rpc-verify_key", AppIdMiyoushe)
	for i := 0; i < len(vs); i++ {
		switch v := vs[i].(type) {
		case config.Device:
			r.SetHeaders(v.Headers())
		case *Validate:
			if v == nil {
				continue
			}
			r.SetHeader("x-rpc-challenge", v.Challenge)
			r.SetHeader("x-rpc-validate", v.Validate)
		}
	}
	return r
}

func Exec[T any](r *resty.Request, method, url string, ds ...int) (t T, _ error) {
	if len(ds) == 0 || ds[0] <= 1 {
		r.SetHeader("DS", DS1())
	} else if ds[0] == 2 {
		r.SetHeader("DS", DS2(r))
	}
	res, err := resty.ParseResp[*BaseResp[any], *BaseResp[T]](
		r.SetError(&BaseResp[any]{}).SetResult(&BaseResp[T]{}).Execute(method, url),
	)
	if err != nil {
		return t, fmt.Errorf("[miyoushe] %w", err)
	}
	return res.Data, nil
}

// DS1 https://github.com/UIGF-org/mihoyo-api-collect/issues/1
func DS1() string {
	_, s := ds1(time.Now().Unix(), alphanum(6))
	return s
}

func ds1(t int64, r string) (string, string) {
	s := fmt.Sprintf("salt=%s&t=%d&r=%s", AppSaltK2, t, r)
	b := md5.Sum([]byte(s))
	return s, fmt.Sprintf("%d,%s,%s", t, r, hex.EncodeToString(b[:]))
}

const dicts = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func alphanum(n int) string {
	bs := make([]byte, n)
	for i := range bs {
		bs[i] = dicts[rand.Intn(len(dicts))]
	}
	return string(bs)
}

func DS2(q *resty.Request) string {
	_, s := ds2(time.Now().Unix(), 100000+rand.Intn(100000), json.MustMarshalString(q.Body), q.QueryParam.Encode())
	return s
}

func ds2(t int64, r int, body, query string) (string, string) {
	if r == 100000 {
		r += 542367
	}
	s := fmt.Sprintf("salt=%s&t=%d&r=%d&b=%s&q=%s", AppSalt6X, t, r, body, query)
	b := md5.Sum([]byte(s))
	return s, fmt.Sprintf("%d,%d,%s", t, r, hex.EncodeToString(b[:]))
}

// generate base stoken auth cookies
func hcSToken(account config.Account) []*http.Cookie {
	return []*http.Cookie{
		{Name: "mid", Value: account.Mid},
		{Name: "stoken", Value: account.SToken},
	}
}

// generate base cookie_token auth cookies
func hcCToken(account config.Account) []*http.Cookie {
	return []*http.Cookie{
		{Name: "account_id", Value: account.Uid},
		{Name: "cookie_token", Value: account.CToken},
	}
}
