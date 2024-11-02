package post

import (
	"time"
)

type ID = uint64

//go:generate accessor -type Post
type Post struct {
	id            ID
	title         string
	likeCount     uint32
	commentsCount uint32
	caption       string
	permalink     string
	postedAt      time.Time
}

func NewPost(
	id ID,
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
