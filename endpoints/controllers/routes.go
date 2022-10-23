package controllers

import (
	"github.com/beijingzhangwei/ddyy-b/endpoints/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/api/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/api/auth/create-user", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/users/{user_id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")

	s.Router.HandleFunc("/users/{user_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{user_id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/api/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/api/one_user_posts/{user_id}", middlewares.SetMiddlewareJSON(s.GetPostsByUserId)).Methods("GET")
	s.Router.HandleFunc("/api/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/api/posts/{post_id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
	s.Router.HandleFunc("/posts/{post_id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{post_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")

	//Posts routes
	s.Router.HandleFunc("/api/posts/{post_id}/comments",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateComments))).Methods("POST")
	s.Router.HandleFunc("/api/comments/{comment_id}/del", middlewares.SetMiddlewareAuthentication(s.DeleteComment)).Methods("POST")

	// 历史版本
	//s.Router.HandleFunc("/api/auth/login", version_v1.GetTokenUserPassword).Methods("POST")
	//s.Router.HandleFunc("/api/auth/create-user", createUser).Methods("POST")
	//s.Router.HandleFunc("/api/posts", version_v1.getPosts).Methods("GET")
	//s.Router.HandleFunc("/api/users/{USERNAME}", version_v1.CheckTokenHandler(version_v1.getUser)).Methods("GET")
	//s.Router.HandleFunc("/api/posts", checkTokenHandler(addPost)).Methods("POST")
	//s.Router.HandleFunc("/api/posts/{POST_ID}", checkTokenHandler(deletePost)).Methods("DELETE")

	//s.Router.HandleFunc("/api/posts/{POST_ID}/comments", checkTokenHandler(addComment)).Methods("POST")

	//s.Router.HandleFunc("/api/auth/token", checkTokenHandler(getTokenByToken)).Methods("GET")
}

type CorsRouterDecorator struct {
	R *mux.Router
}

func (c *CorsRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, PUT, DELETE, PATCH")
		rw.Header().Add("Access-Control-Allow-Headers",
			"Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	c.R.ServeHTTP(rw, req)
}
