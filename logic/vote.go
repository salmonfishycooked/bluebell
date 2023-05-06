package logic

import (
	"go.uber.org/zap"
	"go_web_app/dao/redis"
	"go_web_app/model"
	"strconv"
)

// 投票功能

// VoteForPost 给某条帖子投票
func VoteForPost(userID int64, p *model.VoteData) error {
	zap.L().Debug("VoteForPost", zap.Int64("userID", userID), zap.String("postID", p.PostID), zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
