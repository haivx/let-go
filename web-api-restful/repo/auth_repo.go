package repo

import (
	dto "final-project/dto"
	model "final-project/model"
	services "final-project/services"
)

func Login(currentUser *model.User) (user *dto.LoginResponse, err error) {
	token, err := services.GenerateToken(currentUser)
	if err != nil {
		return nil, err
	}
	user = &dto.LoginResponse{
		Token:       token,
		Username:    currentUser.Username,
		Email:       currentUser.Email,
		PhoneNumber: currentUser.PhoneNumber,
	}
	return user, err
}
