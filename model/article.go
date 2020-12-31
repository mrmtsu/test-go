package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

// Article ...
type Article struct {
	ID      int       `db:"id" form:"id"`
	Title   string    `db:"title" form:"title" validate:"required,max=50"`
	Body    string    `db:"body" form:"body" validate:"required"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}

// ValidationErrors ...
func (a *Article) ValidationErrors(err error) []string {
	var errMessages []string

	for _, err := range err.(validator.ValidationErrors) {
		var message string

		switch err.Field() {
		case "Title":
			switch err.Tag() {
			case "required":
				message = "タイトルは必須です。"
			case "max":
				message = "タイトルは最大50文字です。"
			}
		case "Body":
			message = "本文は必須です。"
		}

		if message != "" {
			errMessages = append(errMessages, message)
		}
	}

	return errMessages
}
