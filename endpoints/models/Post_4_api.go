package models

import "time"

type Post4Api struct {
	ID         uint64         `json:"id"`
	PostID     uint64         `json:"post_id"` // 帖子ID
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	PostAuthor *User4Api      `json:"post_author"`
	Comments   []*Comment4Api `json:"comments"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func TransPost4Api(post *Post, u *User, comment4Apis []*Comment4Api) *Post4Api {
	return &Post4Api{
		ID:         post.ID,
		PostID:     post.PostID,
		Title:      post.Title,
		Content:    post.Content,
		PostAuthor: TransferUserRaw2UserApi(u),
		Comments:   comment4Apis,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
	}
}
