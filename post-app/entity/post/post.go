package post

import (
	"time"
)

type Id = uint64

//go:generate accessor -type Post
type Post struct {
	id            Id
	title         string
	likeCount     uint32
	commentsCount uint32
	caption       string
	permalink     string
	postedAt      time.Time
}

func NewPost(
	id Id,
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
