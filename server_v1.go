package main

import (
	"fmt"
	"github.com/beijingzhangwei/ddyy-b/endpoints/controllers"
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

	http.Handle("/", &controllers.CorsRouterDecorator{R: r})
	fmt.Println("Listening")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
