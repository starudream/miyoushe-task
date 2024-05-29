package ocr

import (
	"fmt"
	"time"

	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/config"
)

func TT(key, gt, challenge, refer string) (*common.Verification, error) {
	resultId, err := ttRecognize(key, gt, challenge, refer)
	if err != nil {
		return nil, err
	}

	vch := make(chan *common.Verification, 1)

	go func() {
		var data *common.Verification
		for {
			time.Sleep(config.TT().Interval)
			data, err = ttResult(key, resultId)
			if e, ok1 := resty.AsRespErr(err); ok1 {
				if v, ok2 := e.Response.Result().(interface{ GetStatus() int }); ok2 && v.GetStatus() == 2 {
					continue
				}
			}
			vch <- data
			return
		}
	}()

	select {
	case v := <-vch:
		return v, err
	case <-time.After(config.TT().Timeout):
		return nil, fmt.Errorf("[ttocr] 识别超时")
	}
}

type ttRecognizeResp struct {
	Status   int    `json:"status"`
	Msg      string `json:"msg"`
	ResultId string `json:"resultid,omitempty"`
}

func (t *ttRecognizeResp) IsSuccess() bool {
	return t.Status == 1
}

func (t *ttRecognizeResp) String() string {
	return fmt.Sprintf("status: %d, msg: %s", t.Status, t.Msg)
}

func ttRecognize(key, gt, challenge, refer string) (string, error) {
	form := gh.MS{"appkey": key, "gt": gt, "challenge": challenge, "itemid": config.TT().ItemId, "referer": refer}
	res, err := resty.ParseResp[*ttRecognizeResp, *ttRecognizeResp](
		resty.R().SetError(&ttRecognizeResp{}).SetResult(&ttRecognizeResp{}).SetFormData(form).Post("https://api.ttocr.com/api/recognize"),
	)
	if err != nil {
		return "", fmt.Errorf("[ttocr] %w", err)
	}
	return res.ResultId, nil
}

type ttResultResp struct {
	Status int                  `json:"status"`
	Msg    string               `json:"msg"`
	Time   int                  `json:"time,omitempty"`
	Data   *common.Verification `json:"data,omitempty"`
}

func (t *ttResultResp) GetStatus() int {
	if t == nil {
		return 0
	}
	return t.Status
}

func (t *ttResultResp) IsSuccess() bool {
	return t.Status == 1
}

func (t *ttResultResp) String() string {
	return fmt.Sprintf("status: %d, msg: %s", t.Status, t.Msg)
}

func ttResult(key, resultId string) (*common.Verification, error) {
	form := gh.MS{"appkey": key, "resultid": resultId}
	res, err := resty.ParseResp[*ttResultResp, *ttResultResp](
		resty.R().SetError(&ttResultResp{}).SetResult(&ttResultResp{}).SetFormData(form).Post("https://api.ttocr.com/api/results"),
	)
	if err != nil {
		return nil, fmt.Errorf("[ttocr] %w", err)
	}
	return res.Data, nil
}
