package formaterror

import (
	"errors"
	"strings"
)

// FormatError 为了以更易读的方式格式化一些错误消息，我们需要创建一个包来帮助我们实现这一点。
func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("Nickname Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	return errors.New(err)
}
