package job

import (
	"fmt"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/api/miyoushe"
	"github.com/starudream/miyoushe-task/api/ocr"
	"github.com/starudream/miyoushe-task/config"
)

func Verify(account config.Account) (verification *common.Verification, _ error) {
	res, err := miyoushe.CreateVerification(account)
	if err != nil {
		return nil, fmt.Errorf("create verification error: %w", err)
	}

	verification, err = DM(res.Gt, res.Challenge)
	if err != nil {
		return nil, err
	}

	_, err = miyoushe.VerifyVerification(verification.Challenge, verification.Validate, account)
	if err != nil {
		return nil, fmt.Errorf("verify verification error: %w", err)
	}

	return verification, nil
}

func DM(gt, challenge string) (*common.Verification, error) {
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

func dm(gt, challenge string) (*common.Verification, error) {
	if key := config.C().RROCRKey; key != "" {
		slog.Info("attempt to dm using rrocr, please wait a moment")
		return ocr.RR(key, gt, challenge, common.RefererAct)
	}
	return nil, nil
}
