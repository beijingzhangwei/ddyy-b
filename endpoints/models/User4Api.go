package models

type User4Api struct {
	ID          uint64 `json:"id"`
	UserID      uint64 `json:"user_id"` // 用户ID
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

func TransferUserRaw2UserApi(user *User) *User4Api {
	return &User4Api{
		ID:          user.ID,
		UserID:      user.UserID,
		Nickname:    user.Nickname,
		Email:       user.Email,
		Description: user.Description,
	}
}
