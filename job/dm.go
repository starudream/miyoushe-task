package job

import (
	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/miyoushe-task/api/miyoushe"
	"github.com/starudream/miyoushe-task/config"
	"github.com/starudream/miyoushe-task/util/ocr"
)

func DM(gt, challenge string) (*miyoushe.Validate, error) {
	if key := config.C().RROCRKey; key != "" {
		slog.Info("attempt to dm using rrocr, please wait a moment")
		return ocr.RR(key, gt, challenge, miyoushe.RefererAct)
	}
	return nil, nil
}
