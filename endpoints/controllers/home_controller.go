package controllers

import (
	responses "github.com/beijingzhangwei/ddyy-b/endpoints/reponses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
