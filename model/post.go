package model

import "time"

type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	Supports   int64  `json:"supports"`

	// SupportStatus 0: none, -1: against, 1: suppose
	SupportStatus int `json:"support_status"`

	*Post                               // 帖子详情
	*CommunityDetail `json:"community"` // 社区详情
}
