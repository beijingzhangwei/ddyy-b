package models

import (
	"errors"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	CommentID uint64    `gorm:"not null;unique" json:"comment_id"` // 评论ID
	PostID    uint64    `gorm:"not null" json:"post_id"`
	AuthorID  uint64    `gorm:"not null" json:"author_id"`
	Username  string    `gorm:"size:100;not null;" json:"username"`
	Content   string    `gorm:"size:255;not null;" json:"content"` // 评论内容
	CreatedAt time.Time `gorm:"column:CreatedAt;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt;autoUpdateTime" json:"updated_at"`
}

func (c *Comment) Prepare() {
	c.Content = html.EscapeString(strings.TrimSpace(c.Content))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.CommentID = uint64(time.Now().Unix())
}

func (c *Comment) Validate() error {
	if c.Content == "" {
		return errors.New("Required Content")
	}
	if c.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

func (c *Comment) SaveComment(db *gorm.DB) (*Comment, error) {
	var err error
	err = db.Debug().Model(&Comment{}).Create(&c).Error
	if err != nil {
		return &Comment{}, err
	}
	return c, nil
}

func (c *Comment) DeleteAComment(db *gorm.DB, commentId uint64, uid uint64) (int64, error) {

	db = db.Debug().Model(&Comment{}).Where("comment_id = ? and author_id = ?", commentId, uid).Take(&Comment{}).Delete(&Comment{})

	if db.Error != nil {
		if gorm.ErrRecordNotFound == db.Error {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (c *Comment) GetCommentsByPostId(db *gorm.DB, postId uint64) ([]*Comment4Api, error) {
	var cs []Comment
	cErr := db.Debug().Model(&Comment{}).Where("post_id in (?)", postId).Limit(100).Find(&cs).Error
	if cErr != nil {
		return nil, cErr
	}
	comment4Apis := make([]*Comment4Api, 0)
	for _, comment := range cs {
		cu := &User{}
		cuErr := db.Debug().Model(&User{}).Where("user_id = ?", comment.AuthorID).Take(cu).Error
		if cuErr != nil {
			return nil, cuErr
		}
		comment4Apis = append(comment4Apis, TransComment4Api(&comment, postId, cu))
	}
	return comment4Apis, nil
}
