package models

import "gopkg.in/validator.v2"

type Users struct {
	Id       int    `json:"id,omitempty"` //Caso seja vazio não pega no Json Marshall
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Email    string `json:"email"`
}

func ValidaDadosClientes(user *Users) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
