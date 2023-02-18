package models

import "errors"

type AiInput struct {
	Input string `json:"ai_input"`
}

func (i AiInput) Validate() error {
	if len(i.Input) == 0 {
		return errors.New("输入不能为空")
	}
	return nil
}
