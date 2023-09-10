package controller

import (
	dto "final-project/dto"
	model "final-project/model"
	mailer "final-project/services/mailer"
	"net/http"

	repo "final-project/repo"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var newUser model.CreateUser

	if err := c.ShouldBindJSON(&newUser); err != nil {
		ResponseError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	data, _ := repo.GetUser(newUser.Email)
	userPass := newUser.Password
	if data != nil {
		ResponseError(c, http.StatusBadRequest, "user's exist")
		return
	}

	encodedHash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)

	newUser.Password = string(encodedHash)

	user, err := repo.CreateUser(&newUser)
	userInfo := mailer.Information{
		Username: user.Username,
		Password: userPass,
	}

	mailer.SendGMail(userInfo)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	ResponseData(c, http.StatusOK, user)
}

func Login(c *gin.Context) {
	var login dto.UserLogin
	if err := c.ShouldBindJSON(&login); err != nil {
		ResponseError(c, http.StatusBadRequest, "Wrong data input")
		return
	}

	dataUser, _ := repo.GetUser(login.Email)
	if dataUser == nil {
		ResponseError(c, http.StatusBadRequest, "user is not exist")
		return
	}

	notValidPassword := bcrypt.CompareHashAndPassword([]byte(string(dataUser.Password)), []byte(login.Password))

	if notValidPassword != nil {
		ResponseError(c, http.StatusBadRequest, "username's not existed or wrong password")
		return
	}

	loginInfo, err := repo.Login(dataUser)
	if err != nil {
		ResponseError(c, http.StatusNotFound, err.Error())
		return
	}

	ResponseData(c, http.StatusOK, loginInfo)
}
