package models

import "time"

type Comment4Api struct {
	ID            int       `json:"id"`
	CommentID     uint64    `json:"comment_id"` // 评论ID
	PostID        uint64    `json:"post_id"`
	CommentAuthor *User4Api `json:"comment_author"`
	Content       string    `json:"content"` // 评论内容
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func TransComment4Api(comment *Comment, postId uint64, cu *User) *Comment4Api {
	return &Comment4Api{
		ID:            comment.ID,
		CommentID:     comment.CommentID,
		PostID:        postId,
		CommentAuthor: TransferUserRaw2UserApi(cu),
		Content:       comment.Content,
		CreatedAt:     comment.CreatedAt,
		UpdatedAt:     comment.UpdatedAt,
	}
}
