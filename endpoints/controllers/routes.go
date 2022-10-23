package controllers

import "github.com/beijingzhangwei/ddyy-b/endpoints/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{user_id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{user_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{user_id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{post_id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{post_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{post_id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/comments", middlewares.SetMiddlewareJSON(s.CreateComments)).Methods("POST")
	s.Router.HandleFunc("/comments/{comment_id}", middlewares.SetMiddlewareAuthentication(s.DeleteComment)).Methods("DELETE")
}

// 	r.HandleFunc("/api/posts", getPosts).Methods("GET")
//	r.HandleFunc("/api/posts", checkTokenHandler(addPost)).Methods("POST")
//	r.HandleFunc("/api/posts/{POST_ID}", checkTokenHandler(deletePost)).Methods("DELETE")
//	r.HandleFunc("/api/posts/{POST_ID}/comments", checkTokenHandler(addComment)).Methods("POST")
//	r.HandleFunc("/api/auth/login", getTokenUserPassword).Methods("POST")
//	r.HandleFunc("/api/auth/create-user", createUser).Methods("POST")
//	r.HandleFunc("/api/auth/token", checkTokenHandler(getTokenByToken)).Methods("GET")
//	r.HandleFunc("/api/users/{USERNAME}", checkTokenHandler(getUser)).Methods("GET")
