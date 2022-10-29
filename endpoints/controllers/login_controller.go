package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beijingzhangwei/ddyy-b/endpoints/auth"
	"github.com/beijingzhangwei/ddyy-b/endpoints/models"
	"github.com/beijingzhangwei/ddyy-b/endpoints/reponses"
	"github.com/beijingzhangwei/ddyy-b/endpoints/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dbUser, token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, struct {
		Token  string `json:"token"`
		Email  string `json:"email"`
		UserId uint64 `json:"user_id"`
	}{
		Token:  token,
		Email:  dbUser.Email,
		UserId: dbUser.UserID,
	})
}

func (server *Server) SignIn(email, password string) (*models.User, string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		fmt.Println("SignIn query user error：", err)
		return nil, "", err
	}
	err = models.VerifyPassword(user.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("SignIn.VerifyPassword user.Password：", user.Password, "password", password)
		fmt.Println("SignIn.VerifyPassword error：", err)
		return nil, "", err
	}
	token, err := auth.CreateToken(user.UserID)
	return &user, token, err
}
