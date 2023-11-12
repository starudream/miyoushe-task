package job

import (
	"bytes"
	"fmt"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/miyoushe-task/api/miyoushe"
	"github.com/starudream/miyoushe-task/config"
)

func SignBBS(gameId string, account config.Account) (err error) {
	if gameId == "" {
		gameId = miyoushe.GameIdSR
	}

	account, err = Refresh(account)
	if err != nil {
		return err
	}

	game := miyoushe.GameById[gameId]

	slog.Info("attempt to sign bbs %s", game.Name)

	var validate *miyoushe.Validate

sign:

	res, err := miyoushe.SignBBS(gameId, account, validate)
	if err != nil {
		if miyoushe.IsRetCode(err, miyoushe.RetCodeBBSHasSigned) {
			slog.Info("bbs has signed")
			return nil
		} else if miyoushe.IsRetCode(err, miyoushe.RetCodeBBSNeedValidate) {
			validate, err = Verify(account)
			if err != nil {
				return fmt.Errorf("verify error: %w", err)
			}
			goto sign
		}
		return fmt.Errorf("sign bbs error: %w", err)
	}

	slog.Info("sign bbs %s success and got %d points", game.Name, res.Points)
	return nil
}

func Verify(account config.Account) (validate *miyoushe.Validate, _ error) {
	res, err := miyoushe.CreateVerification(account)
	if err != nil {
		return nil, fmt.Errorf("create verification error: %w", err)
	}

	validate, err = DM(res.Gt, res.Challenge)
	if err != nil {
		return nil, fmt.Errorf("dm error: %w", err)
	}
	if validate == nil {
		return nil, fmt.Errorf("verify maybe risk")
	}

	_, err = miyoushe.VerifyVerification(validate.Challenge, validate.Validate, account)
	if err != nil {
		return nil, fmt.Errorf("verify verification error: %w", err)
	}

	return validate, nil
}

func SignLuna(account config.Account) (_ map[string]map[string]string, err error) {
	account, err = Refresh(account)
	if err != nil {
		return
	}

	res, err := miyoushe.ListGameRole("", account)
	if err != nil {
		return nil, fmt.Errorf("list game role error: %w", err)
	}
	if len(res.List) == 0 {
		return nil, fmt.Errorf("no binding game roles")
	}

	return SignLunaGame(res.List, account)
}

func SignLunaGame(roles []*miyoushe.GameRole, account config.Account) (map[string]map[string]string, error) {
	awards := map[string]map[string]string{} // key: game name, value: map[key: role nickname+uid, value: awards]

	for _, role := range roles {
		gameId, exists := miyoushe.GameIdByBiz[role.GameBiz]
		if !exists {
			slog.Warn("game biz %s not supported", role.GameBiz)
			continue
		}

		game := miyoushe.GameById[gameId]

		actId, err := GetLunaActId(gameId)
		if err != nil {
			return nil, err
		}

		if _, ok := awards[game.Name]; !ok {
			awards[game.Name] = map[string]string{}
		}

		slog.Info("attempt to sign luna %s %s (%s)", game.Name, role.Nickname, role.GameUid)

		var validate *miyoushe.Validate

		retry := config.C().DMRetry

	sign:

		res1, err := miyoushe.SignLuna(actId, role.Region, role.GameUid, account, validate)
		if err != nil {
			if !miyoushe.IsRetCode(err, miyoushe.RetCodeLunaHasSigned) {
				return nil, fmt.Errorf("sign luna error: %w", err)
			}
			slog.Info("luna has signed")
		} else if res1.IsRisky() {
			slog.Warn("sign luna maybe risk, gt: %s, challenge: %s", res1.Gt, res1.Challenge)
			if retry < 0 {
				return nil, fmt.Errorf("dm retry limit exceeded")
			}
			slog.Info("attempt to dm bypass verification")
			validate, err = DM(res1.Gt, res1.Challenge)
			if err != nil {
				return nil, fmt.Errorf("dm error: %w", err)
			}
			if validate == nil {
				return nil, fmt.Errorf("sign luna maybe risk")
			}
			retry--
			goto sign
		}

		res2, err := miyoushe.ListLunaAward(actId, role.Region, role.GameUid, account)
		if err != nil {
			return nil, fmt.Errorf("list luna award error: %w", err)
		}

		s1, s2 := fmt.Sprintf("%s (%s)", role.Nickname, role.GameUid), res2.Today().ShortString()
		awards[game.Name][s1] = s2
		slog.Info("sign luna %s %s success and got %s", game.Name, s1, s2)
	}

	return awards, nil
}

func GetLunaActId(gameId string) (string, error) {
	res, err := miyoushe.GetBBSHome(gameId)
	if err != nil {
		return "", fmt.Errorf("get bbs home error: %w", err)
	}
	actId := res.GetLunaActId()
	if actId == "" {
		return "", fmt.Errorf("luna act id not found")
	}
	return actId, nil
}

func FormatAwards(awards map[string]map[string]string) string {
	buf := &bytes.Buffer{}
	for game, roles := range awards {
		buf.WriteString("【")
		buf.WriteString(game)
		buf.WriteString("】\n")
		for name, award := range roles {
			buf.WriteString("  ")
			buf.WriteString(name)
			buf.WriteString(" ")
			if award != "" {
				buf.WriteString(award)
			} else {
				buf.WriteString("失败")
			}
			buf.WriteString("\n")
		}
	}
	return buf.String()
}
