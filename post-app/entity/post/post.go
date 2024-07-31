package post

import (
	"time"
)

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

func (p *Post) GetId() uint64 {
	return p.id
}

func (p *Post) GetTitle() string {
	return p.title
}

func (p *Post) GetLikeCount() uint32 {
	return p.likeCount
}

func (p *Post) GetCommentsCount() uint32 {
	return p.commentsCount
}

func (p *Post) GetCaption() string {
	return p.caption
}

func (p *Post) GetPermalink() string {
	return p.permalink
}

func (p *Post) GetPostedAt() time.Time {
	return p.postedAt
}
