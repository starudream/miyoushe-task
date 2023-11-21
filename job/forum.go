package job

import (
	"fmt"
	"strings"
	"time"

	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/miyoushe-task/api/common"
	"github.com/starudream/miyoushe-task/api/miyoushe"
	"github.com/starudream/miyoushe-task/config"
)

const (
	VerifyRetry = 3

	PostView   = 3
	PostUpvote = 10
	PostShare  = 1
	PostLoop   = 10
)

type SignForumRecord struct {
	GameId    string
	GameName  string
	HasSigned bool
	IsRisky   bool
	Verify    int
	Points    int

	PostView   int
	PostUpvote int
	PostShare  int
	LoopCount  int
}

func (r SignForumRecord) Name() string {
	return "米游社每日任务"
}

func (r SignForumRecord) Success() string {
	vs := []string{r.Name() + "完成"}
	if r.IsRisky {
		vs = append(vs, "使用打码平台验证")
	}
	vs = append(vs, fmt.Sprintf("在版区【%s】", r.GameName))
	if r.Points > 0 {
		vs = append(vs, fmt.Sprintf(" 打卡获得%d米游币", r.Points))
	}
	vs = append(vs,
		fmt.Sprintf(" 浏览%d/%d个帖子", r.PostView, PostView),
		fmt.Sprintf(" 点赞%d/%d个帖子", r.PostUpvote, PostUpvote),
		fmt.Sprintf(" 分享%d/%d个帖子", r.PostShare, PostShare),
	)
	return strings.Join(vs, "\n")
}

func SignForum(account config.Account) (record SignForumRecord, err error) {
	account, err = Refresh(account)
	if err != nil {
		return
	}

	businesses, err := miyoushe.GetBusinesses(account)
	if err != nil {
		err = fmt.Errorf("get businesses error: %w", err)
		return
	}

	if len(businesses.Businesses) == 0 {
		err = fmt.Errorf("no channel subscription, please use phone to login and subscribe channel")
		return
	}

	defer func() {
		slog.Info("sign forum record: %+v", record)
		if err != nil {
			slog.Error("sign forum error: %v", err)
		}
	}()

	gameId := businesses.Businesses[0]
	game := miyoushe.AllGamesById[gameId]

	record.GameId = gameId
	record.GameName = game.Name

	today, err := miyoushe.GetSignForum(gameId, account)
	if err != nil {
		err = fmt.Errorf("get sign forum error: %w", err)
		return
	}

	var (
		verification  *common.Verification
		signForumData *miyoushe.SignForumData
	)

	if today.IsSigned {
		record.HasSigned = true
		goto post
	}

sign:

	signForumData, err = miyoushe.SignForum(gameId, account, verification)
	if err != nil {
		if common.IsRetCode(err, common.RetCodeForumHasSigned) {
			record.HasSigned = true
		} else if common.IsRetCode(err, common.RetCodeForumNeedVerification) {
			record.IsRisky = true
		verify:
			record.Verify++
			verification, err = Verify(account)
			if err != nil {
				slog.Error("verify error: %v", err)
				if record.Verify < VerifyRetry {
					slog.Info("retry verify, count: %d", record.Verify)
					goto verify
				}
			}
			goto sign
		} else {
			err = fmt.Errorf("sign forum error: %w", err)
			return
		}
	}

	record.Points = signForumData.Points

post:

	record.LoopCount++

	posts, err := miyoushe.ListFeedPost(gameId, account)
	if err != nil {
		err = fmt.Errorf("list feed post error: %w", err)
		return
	}

	for i := 0; i < len(posts.List); i++ {
		p := posts.List[i]
		pid := p.Post.PostId
		if record.PostView < PostView {
			_, e := miyoushe.GetPost(pid, account)
			if e != nil {
				slog.Error("get post error: %v", e)
				continue
			}
			record.PostView++
		}
		if record.PostUpvote < PostUpvote && !p.IsUpvote() {
			e := miyoushe.UpvotePost(pid, false, account)
			if e != nil {
				slog.Error("upvote post error: %v", e)
				continue
			}
			slog.Debug("upvote post: %s (%s) %s", p.Post.Subject, pid, p.User.Nickname)
			record.PostUpvote++
		}
		if record.PostShare < PostShare {
			_, e := miyoushe.SharePost(pid, account)
			if e != nil {
				slog.Error("share post error: %v", e)
				continue
			}
			record.PostShare++
		}
		// avoid too fast
		time.Sleep(500 * time.Millisecond)
	}

	if record.LoopCount < PostLoop && (record.PostView < PostView || record.PostUpvote < PostUpvote || record.PostShare < PostShare) {
		goto post
	}

	return
}
