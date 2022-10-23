package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	PostID    uint64    `gorm:"not null" json:"post_id"` // 帖子ID
	Title     string    `gorm:"size:100;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	AuthorID  uint64    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"column:CreatedAt;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt;autoUpdateTime" json:"updated_at"`
}

func (p *Post) Prepare() {
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

func (p *Post) SavePost(db *gorm.DB) (*Post4Api, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post4Api{}, err
	}
	u := &User{}
	err = db.Debug().Model(&User{}).Where("user_id = ?", p.AuthorID).Take(u).Error
	if err != nil {
		return nil, err
	}
	return TransPost4Api(p, u, nil), nil
}

func (p *Post) FindAllPosts(db *gorm.DB, uid uint64) ([]*Post4Api, error) {
	var err error
	posts := []Post{}
	if uid == 0 {
		err = db.Debug().Model(&Post{}).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: clause.PrimaryKey},
			Desc:   true,
		}).Limit(100).Find(&posts).Error
	} else {
		err = db.Debug().Model(&Post{}).Where("author_id = ?", uid).Order(clause.OrderByColumn{
			Column: clause.Column{Table: clause.CurrentTable, Name: clause.PrimaryKey},
			Desc:   true,
		}).Limit(100).Find(&posts).Error
	}

	if err != nil {
		return nil, err
	}

	post4Apis := make([]*Post4Api, 0)
	// 后面可以优化 不要循环查询 -- batch
	if len(posts) > 0 {
		for _, post := range posts {
			// 作者
			u := &User{}
			err := db.Debug().Model(&User{}).Where("user_id = ?", post.AuthorID).Take(u).Error
			if err != nil {
				return nil, err
			}
			// 评论
			var cs []Comment
			cErr := db.Debug().Model(&Comment{}).Where("post_id in (?)", post.PostID).Limit(100).Find(&cs).Error
			if cErr != nil {
				return nil, cErr
			}
			// 评论的作者查询
			comment4Apis := make([]*Comment4Api, 0)
			for _, comment := range cs {
				cu := &User{}
				cuErr := db.Debug().Model(&User{}).Where("user_id = ?", comment.AuthorID).Take(cu).Error
				if cuErr != nil {
					return nil, cuErr
				}
				comment4Apis = append(comment4Apis, TransComment4Api(&comment, post.PostID, cu))
			}
			post4Apis = append(post4Apis, TransPost4Api(&post, u, comment4Apis))
		}
	}
	return post4Apis, nil
}

func (p *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("post_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}
	return p, nil
}

func (p *Post) UpdateAPost(db *gorm.DB) (*Post, error) {

	err := db.Debug().Model(&Post{}).Where("post_id = ?", p.PostID).Updates(Post{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
	fmt.Println("UpdateAPost.Updates err：", err)
	if err != nil {
		return &Post{}, err
	}
	return p, nil
}

func (p *Post) UpdateAPostWithTx(db *gorm.DB) (*Post, error) {

	errTx := db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		err := db.Debug().Model(&Post{}).Where("post_id = ?", p.PostID).Updates(Post{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
		fmt.Println("UpdateAPost.Updates err：", err)
		if err != nil {
			return err
		}
		return nil
	})
	if errTx != nil {
		return &Post{}, errTx
	}
	return p, nil
}

func (p *Post) DeleteAPost(db *gorm.DB, pid uint64, uid uint64) (int64, error) {

	db = db.Debug().Model(&Post{}).Where("post_id = ? and author_id = ?", pid, uid).Take(&Post{}).Delete(&Post{})

	if db.Error != nil {
		if gorm.ErrRecordNotFound == db.Error {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
