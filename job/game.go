package job

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/api/miyoushe"
	"github.com/starudream/miyoushe-task/config"
)

var SignGameIdByBiz = map[string]string{
	common.GameBizYSCN: common.GameIdYS,
	common.GameBizSRCN: common.GameIdSR,
}

type SignGameRecord struct {
	GameId    string
	GameName  string
	RoleName  string
	RoleUid   string
	HasSigned bool
	IsRisky   bool
	IsSuccess bool
	Verify    int
	Award     string
}

type SignGameRecords []SignGameRecord

func (rs SignGameRecords) Name() string {
	return "米游社游戏签到"
}

func (rs SignGameRecords) Success() string {
	vs := []string{rs.Name() + "完成"}
	for i := 0; i < len(rs); i++ {
		if rs[i].HasSigned || rs[i].IsSuccess {
			vs = append(vs, fmt.Sprintf("在游戏【%s】角色【%s】获得 %s（%d）", rs[i].GameName, rs[i].RoleName, rs[i].Award, rs[i].Verify))
		}
	}
	return strings.Join(vs, "\n")
}

func SignGame(account config.Account) (_ SignGameRecords, err error) {
	account, err = RefreshSTokenAuto(account)
	if err != nil {
		return nil, err
	}
	account, err = RefreshCToken(account)
	if err != nil {
		return nil, err
	}
	roles, err := miyoushe.ListGameRole("", account)
	if err != nil {
		return nil, fmt.Errorf("list game role error: %w", err)
	}
	return SignGameRoles(roles.List, account)
}

func SignGameRoles(roles []*miyoushe.GameRole, account config.Account) (SignGameRecords, error) {
	var records []SignGameRecord
	for _, role := range roles {
		record, err := SignGameRole(role, account)
		slog.Info("sign game record: %+v", record)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	slices.SortFunc(records, func(a, b SignGameRecord) int {
		if a.GameId == b.GameId {
			return cmp.Compare(a.RoleUid, b.RoleUid)
		}
		return cmp.Compare(a.GameId, b.GameId)
	})
	return records, nil
}

func SignGameRole(role *miyoushe.GameRole, account config.Account) (record SignGameRecord, err error) {
	record.RoleName = role.Nickname
	record.RoleUid = role.GameUid

	gameId := SignGameIdByBiz[role.GameBiz]
	if gameId == "" {
		slog.Warn("game biz %s not supported", role.GameBiz)
		return
	}

	game := miyoushe.AllGamesById[gameId]

	record.GameId = gameId
	record.GameName = game.Name

	gameName := strings.Split(role.GameBiz, "_")[0]

	home, err := miyoushe.GetHome(gameId)
	if err != nil {
		err = fmt.Errorf("get home error: %w", err)
		return
	}

	actId := home.GetSignActId()
	if actId == "" {
		err = fmt.Errorf("get sign act id error: %w", err)
		return
	}

	today, err := miyoushe.GetSignGame(gameName, actId, role.Region, role.GameUid, account)
	if err != nil {
		err = fmt.Errorf("get sign game error: %w", err)
		return
	}

	var (
		verification *common.Verification
		signGameData *miyoushe.SignGameData
	)

	if today.IsSign {
		record.HasSigned = true
		goto award
	}

sign:

	signGameData, err = miyoushe.SignGame(gameName, actId, role.Region, role.GameUid, account, verification)
	if err != nil {
		if common.IsRetCode(err, common.RetCodeGameHasSigned) {
			record.HasSigned = true
		} else {
			err = fmt.Errorf("sign game error: %w", err)
			return
		}
	} else if signGameData.IsRisky() {
		record.IsRisky = true
		if signGameData.Gt == "" || signGameData.Challenge == "" {
			err = fmt.Errorf("sign game is risky, but gt or challenge is empty")
			return
		}
		record.Verify++
		verification, err = DM(signGameData.Gt, signGameData.Challenge)
		if err != nil {
			slog.Error("dm error: %v", err)
			if record.Verify >= VerifyRetry {
				return
			}
			slog.Info("retry sign, count: %d", record.Verify)
		}
		goto sign
	} else {
		record.IsSuccess = true
	}

award:

	award, err := miyoushe.ListSignGameAward(gameName, actId, role.Region, role.GameUid, account)
	if err != nil {
		err = fmt.Errorf("list sign game award error: %w", err)
		return
	}
	record.Award = award.Today().ShortString()
	return
}
