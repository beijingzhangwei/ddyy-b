package controllers

import (
	"encoding/json"
	"github.com/beijingzhangwei/ddyy-b/endpoints/models"
	responses "github.com/beijingzhangwei/ddyy-b/endpoints/reponses"
	"io/ioutil"
	"net/http"
)

func (server *Server) AI(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	aiInput := models.AiInput{}
	err = json.Unmarshal(body, &aiInput)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = aiInput.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//c := gogpt.NewClient("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")
	//ctx := context.Background()
	//
	//req := gogpt.CompletionRequest{
	//	Model:     gogpt.GPT3TextDavinci003,
	//	MaxTokens: 4000,
	//	Prompt:    aiInput.Input,
	//}
	key1 := "sk-"
	key2 := "YpRB3nYVVwdYlNEHFN5rT3B"
	key3 := "lbkFJnHe2jZXHTuLTqTRzibH7"
	//resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		//responses.ERROR(w, http.StatusInternalServerError, errors.New("ai err"+err.Error()))
		responses.JSON(w, http.StatusOK, struct {
			Answer string `json:"answer"`
		}{
			Answer: "" + key1 + key2 + key3,
		})
		return
	}
	responses.JSON(w, http.StatusOK, struct {
		Answer string `json:"answer"`
	}{
		Answer: "" + key1 + key2 + key3,
	})
}
