// Code generated by "accessor -type Post"; DO NOT EDIT.

package post

import "time"

// Id return id value
func (t Post) Id() uint64 {
	return t.id
}

// Title return title value
func (t Post) Title() string {
	return t.title
}

// LikeCount return likeCount value
func (t Post) LikeCount() uint32 {
	return t.likeCount
}

// CommentsCount return commentsCount value
func (t Post) CommentsCount() uint32 {
	return t.commentsCount
}

// Caption return caption value
func (t Post) Caption() string {
	return t.caption
}

// Permalink return permalink value
func (t Post) Permalink() string {
	return t.permalink
}

// PostedAt return postedAt value
func (t Post) PostedAt() time.Time {
	return t.postedAt
}
