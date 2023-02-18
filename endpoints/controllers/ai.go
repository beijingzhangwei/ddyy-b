package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/beijingzhangwei/ddyy-b/endpoints/models"
	responses "github.com/beijingzhangwei/ddyy-b/endpoints/reponses"
	gogpt "github.com/sashabaranov/go-gpt3"
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

	c := gogpt.NewClient("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")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 4000,
		Prompt:    aiInput.Input,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("ai err"+err.Error()))
		return
	}
	responses.JSON(w, http.StatusOK, struct {
		Answer string `json:"answer"`
	}{
		Answer: resp.Choices[0].Text,
	})
}
