package common

import (
	"fmt"
	"net/http"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/miyoushe-task/config"
)

const (
	UserAgent = "Mozilla/5.0 (Linux; Android 13; 22011211C Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/104.0.5112.97 Mobile Safari/537.36 miHoYoBBS/" + AppVersion

	RefererApp = "https://app.mihoyo.com"
	RefererAct = "https://act.mihoyo.com"

	AppVersion    = "2.81.1"
	AppIdMiyoushe = "bll8iq97cem8" // 米游社
)

func R(vs ...any) *resty.Request {
	r := resty.R().
		SetHeader("Accept-Encoding", "gzip").
		SetHeader("User-Agent", UserAgent).
		SetHeader("Referer", RefererApp).
		SetHeader("x-rpc-app_version", AppVersion).
		SetHeader("x-rpc-app_id", AppIdMiyoushe).
		SetHeader("x-rpc-verify_key", AppIdMiyoushe)
	for i := 0; i < len(vs); i++ {
		switch v := vs[i].(type) {
		case config.Device:
			r.SetHeaders(v.Headers())
		case *Verification:
			if v == nil || v.Challenge == "" || v.Validate == "" {
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

func GetHeaders(err error) http.Header {
	if err != nil {
		e, ok := resty.AsRespErr(err)
		if ok {
			return e.Response.Header()
		}
	}
	return http.Header{}
}

type Aigis struct {
	SessionId string `json:"session_id"`
	MmtType   int    `json:"mmt_type"`
	Data      string `json:"data"`
}

type AigisData struct {
	Challenge  string `json:"challenge"`
	Gt         string `json:"gt"`
	NewCaptcha int    `json:"new_captcha"`
	Success    int    `json:"success"`
}

func GetAigisData(err error) (*Aigis, *AigisData) {
	s := GetHeaders(err).Get("x-rpc-aigis")
	if s != "" {
		a, e1 := json.UnmarshalTo[*Aigis](s)
		if e1 == nil {
			b, e2 := json.UnmarshalTo[*AigisData](a.Data)
			if e2 == nil {
				return a, b
			}
		}
	}
	return nil, nil
}
