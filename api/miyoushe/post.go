package miyoushe

import (
	"github.com/starudream/go-lib/core/v2/gh"

	"github.com/starudream/miyoushe-task/config"
)

type ListPostData struct {
	IsLast   bool    `json:"is_last"`
	IsOrigin bool    `json:"is_origin"`
	LastId   string  `json:"last_id"`
	List     []*Post `json:"list"`
}

type Post struct {
	Post          *PostInfo          `json:"post"`
	Stat          *PostStat          `json:"stat"`
	User          *PostUser          `json:"user"`
	SelfOperation *PostSelfOperation `json:"self_operation"`
}

func (p *Post) IsUpvote() bool {
	return p != nil && p.SelfOperation != nil && p.SelfOperation.Attitude == 1
}

func (p *Post) IsCollected() bool {
	return p != nil && p.SelfOperation != nil && p.SelfOperation.IsCollected
}

type PostInfo struct {
	PostId    string `json:"post_id"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
	MaxFloor  int    `json:"max_floor"`
	ReplyTime string `json:"reply_time"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type PostStat struct {
	BookmarkNum int `json:"bookmark_num"`
	ForwardNum  int `json:"forward_num"`
	LikeNum     int `json:"like_num"`
	ReplyNum    int `json:"reply_num"`
	ViewNum     int `json:"view_num"`
}

type PostUser struct {
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
}

type PostSelfOperation struct {
	Attitude    int  `json:"attitude"`
	IsCollected bool `json:"is_collected"`
	UpvoteType  int  `json:"upvote_type"`
}

func ListPost(forumId, lastId string, account config.Account) (*ListPostData, error) {
	query := gh.MS{"forum_id": forumId, "is_good": "false", "is_hot": "false", "sort_type": "1", "last_id": lastId, "page_size": "10"}
	return Exec[*ListPostData](R(account.Device).SetCookies(hcSToken(account)).SetQueryParams(query), "GET", AddrTakumi+"/post/api/getForumPostList")
}

type GetPostData struct {
	Post *Post `json:"post"`
}

func GetPost(postId string, account config.Account) (*GetPostData, error) {
	return Exec[*GetPostData](R(account.Device).SetCookies(hcSToken(account)).SetQueryParam("post_id", postId), "GET", AddrTakumi+"/post/api/getPostFull")
}

func UpvotePost(postId string, account config.Account) error {
	_, err := Exec[any](R(account.Device).SetCookies(hcSToken(account)).SetBody(gh.M{"post_id": postId, "is_cancel": false}), "POST", AddrTakumi+"/apihub/sapi/upvotePost")
	return err
}

type SharePostData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Icon    string `json:"icon"`
	Url     string `json:"url"`
}

func SharePost(postId string, account config.Account) (*SharePostData, error) {
	return Exec[*SharePostData](R(account.Device).SetCookies(hcSToken(account)).SetQueryParam("entity_type", "1").SetQueryParam("entity_id", postId), "GET", AddrTakumi+"/apihub/api/getShareConf")
}
