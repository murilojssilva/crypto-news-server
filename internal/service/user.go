package service

import (
	"crypto-news-server/internal/models"
	"errors"
)



func ValidateUserEmail(user *models.User, existingUsers []models.User) error {
	if user.Email == "" {
		return errors.New("endereço de e-mail não foi inserido ou está vazio")
	}

	for _, u := range existingUsers {
		if u.Email == user.Email && u.ID != user.ID {
			return errors.New("já há usuário com esse endereço de e-mail")
		}
	}
	return nil
}


func ValidateUserPassword(user *models.User) error {
	if user.Password == "" {
		return errors.New("senha não foi inserida ou está vazia")
	}
	return nil
}