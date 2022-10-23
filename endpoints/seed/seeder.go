package seed

import (
	"github.com/beijingzhangwei/ddyy-b/endpoints/models"
	"gorm.io/gorm"
	"log"
)

// 该步骤可选 ：可以在添加真实数据之前向数据库添加虚拟数据。
// 本项目中不做实现

var users = []models.User{
	models.User{
		UserID:   100,
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		UserID:   101,
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		PostID:  10,
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		PostID:  11,
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

var comments = []models.Comment{
	models.Comment{
		CommentID: 1000,
		Content:   "Comment --->Hello world 1",
	},
	models.Comment{
		CommentID: 1001,
		Content:   "Comment --->Hello world 2",
	},
}

func Load(db *gorm.DB) {

	//db := s.clone()
	//	for _, value := range values {
	//		if s.HasTable(value) {
	//			db.AddError(s.DropTable(value).Error)
	//		}
	//	}
	//	return db
	// var (
	//		scope     = s.NewScope(value)
	//		tableName string
	//	)
	//
	//	if name, ok := value.(string); ok {
	//		tableName = name
	//	} else {
	//		tableName = scope.TableName()
	//	}
	//
	//	has := scope.Dialect().HasTable(tableName)
	//	s.AddError(scope.db.Error)
	//	return has

	//err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	//if err != nil {
	//	log.Fatalf("cannot drop table: %v", err)
	//}
	//err := db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}).Error
	//if err != nil {
	//	log.Fatalf("cannot migrate table: %v", err)
	//}

	//err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	//if err != nil {
	//	log.Fatalf("attaching foreign key error: %v", err)
	//}

	for i, _ := range users {
		err := db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].UserID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}

		comments[i].AuthorID = users[i].UserID
		comments[i].PostID = posts[i].PostID
		err = db.Debug().Model(&models.Comment{}).Create(&comments[i]).Error
		if err != nil {
			log.Fatalf("cannot seed comments table: %v", err)
		}
	}
}
