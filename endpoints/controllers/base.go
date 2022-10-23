package controllers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

// Server 这个文件将有我们的数据库连接信息，初始化我们的路线，并启动我们的服务器:
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}

	}
	//server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}) //database migration
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 3000")
	log.Fatal(http.ListenAndServe(addr, &CorsRouterDecorator{R: server.Router})) // 支持跨域
}

//
//func mainssss() {
//	r := mux.NewRouter()
//	r = endpoints.AddRouterEndpoints(r)
//	fs := http.FileServer(http.Dir("./dist"))
//	r.PathPrefix("/").Handler(fs)
//
//	http.Handle("/", &corsRouterDecorator{r})
//	fmt.Println("Listening")
//	log.Panic(
//		http.ListenAndServe(":3000", nil),
//	)
//}
//
//type corsRouterDecorator struct {
//	R *mux.Router
//}
//
//func (c *corsRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
//	if origin := req.Header.Get("Origin"); origin != "" {
//		rw.Header().Set("Access-Control-Allow-Origin", origin)
//		rw.Header().Set("Access-Control-Allow-Methods",
//			"POST, GET, PUT, DELETE, PATCH")
//		rw.Header().Add("Access-Control-Allow-Headers",
//			"Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
//	}
//	// Stop here if its Preflighted OPTIONS request
//	if req.Method == "OPTIONS" {
//		return
//	}
//	c.R.ServeHTTP(rw, req)
//}
