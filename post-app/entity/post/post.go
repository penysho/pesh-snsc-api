package post

import (
	"time"
)

//go:generate accessor -type Post
type Post struct {
	id            uint64
	title         string
	likeCount     uint32
	commentsCount uint32
	caption       string
	permalink     string
	postedAt      time.Time
}

func NewPost(
	id uint64,
	title string,
	likeCount uint32,
	commentsCount uint32,
	caption string,
	permalink string,
	postedAt time.Time,
) *Post {
	return &Post{
		id:            id,
		title:         title,
		likeCount:     likeCount,
		commentsCount: commentsCount,
		caption:       caption,
		permalink:     permalink,
		postedAt:      postedAt,
	}
}
