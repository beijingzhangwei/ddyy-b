package main

import (
	"fmt"
	"github.com/beijingzhangwei/ddyy-b/endpoints/version_v1"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main_() {
	r := mux.NewRouter()
	r = version_v1.AddRouterEndpoints(r)
	fs := http.FileServer(http.Dir("./dist"))
	r.PathPrefix("/").Handler(fs)

	http.Handle("/", &corsRouterDecorator{r})
	fmt.Println("Listening")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

type corsRouterDecorator struct {
	R *mux.Router
}

func (c *corsRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
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
