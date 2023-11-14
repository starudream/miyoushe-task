package common

import (
	"fmt"

	"github.com/starudream/go-lib/resty/v2"
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
