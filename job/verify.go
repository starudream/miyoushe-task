package job

import (
	"fmt"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/miyoushe-task/api/miyoushe"
	"github.com/starudream/miyoushe-task/config"
	"github.com/starudream/miyoushe-task/util/ocr"
)

func Verify(account config.Account) (validate *miyoushe.Validate, _ error) {
	res, err := miyoushe.CreateVerification(account)
	if err != nil {
		return nil, fmt.Errorf("create verification error: %w", err)
	}

	validate, err = DM(res.Gt, res.Challenge)
	if err != nil {
		return nil, err
	}

	_, err = miyoushe.VerifyVerification(validate.Challenge, validate.Validate, account)
	if err != nil {
		return nil, fmt.Errorf("verify verification error: %w", err)
	}

	return validate, nil
}

func DM(gt, challenge string) (*miyoushe.Validate, error) {
	validate, err := dm(gt, challenge)
	if err != nil {
		return nil, fmt.Errorf("dm error: %w", err)
	}
	if validate == nil {
		return nil, fmt.Errorf("dm is not configured")
	}
	slog.Info("verification code has been sent to dm and verify success")
	return validate, nil
}

func dm(gt, challenge string) (*miyoushe.Validate, error) {
	if key := config.C().RROCRKey; key != "" {
		slog.Info("attempt to dm using rrocr, please wait a moment")
		return ocr.RR(key, gt, challenge, miyoushe.RefererAct)
	}
	return nil, nil
}
