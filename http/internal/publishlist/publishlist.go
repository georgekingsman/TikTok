package publishlist

import "context"

type GetPublishListService interface {
	GetPublishList(context.Context, GetPublishListReq) (GetPublishListResp, error)
}

type GetPublishListReq struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type GetPublishListResp struct {
	StatusCode int32    `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	VideoList  []*Video `json:"video_list"`
}

type Video struct {
	ID            int64  `json:"id"`
	Author        *User  `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
	PublishTime   int64
}

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}
