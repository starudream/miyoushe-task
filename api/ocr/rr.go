package ocr

import (
	"fmt"

	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/miyoushe-task/api/common"
)

type RRResp struct {
	Status int                  `json:"status"`
	Msg    string               `json:"msg"`
	Code   int                  `json:"code,omitempty"`
	Time   int                  `json:"time,omitempty"`
	Data   *common.Verification `json:"data,omitempty"`
}

func (t *RRResp) IsSuccess() bool {
	return t.Status == 0
}

func (t *RRResp) String() string {
	return fmt.Sprintf("status: %d, msg: %s, code: %d", t.Status, t.Msg, t.Code)
}

func RR(key, gt, challenge, refer string) (*common.Verification, error) {
	form := gh.MS{"appkey": key, "gt": gt, "challenge": challenge, "referer": refer}
	res, err := resty.ParseResp[*RRResp, *RRResp](
		resty.R().SetError(&RRResp{}).SetResult(&RRResp{}).SetFormData(form).Post("http://api.rrocr.com/api/recognize.html"),
	)
	if err != nil {
		return nil, fmt.Errorf("[rrocr] %w", err)
	}
	return res.Data, nil
}
