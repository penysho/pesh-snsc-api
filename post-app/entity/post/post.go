package post

import "time"

type Post struct {
	id            int
	title         string
	likeCount     int
	commentsCount int
	caption       string
	permalink     string
	postedAt      time.Time
}

func NewPost(
	id int,
	title string,
	likeCount int,
	commentsCount int,
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

func (p *Post) GetId() int {
	return p.id
}

func (p *Post) GetTitle() string {
	return p.title
}

func (p *Post) GetLikeCount() int {
	return p.likeCount
}

func (p *Post) GetCommentsCount() int {
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
