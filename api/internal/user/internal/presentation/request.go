package http

import "errors"

type RegisterRequest struct {
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

func (r RegisterRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	if r.PasswordConfirmation == "" {
		return errors.New("password confirmation is required")
	}
	if r.Password != r.PasswordConfirmation {
		return errors.New("passwords do not match")
	}
	return nil
}
