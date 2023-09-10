package controller

import (
	"final-project/dto"
	repo "final-project/repo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	email := c.Query("email")

	user, err := repo.GetUser(email)

	if err != nil {
		ResponseData(c, http.StatusOK, nil)
		return
	}
	ResponseData(c, http.StatusOK, user)
}

func UpdateUserRole(c *gin.Context) {

	var updateUserRole dto.UpdateUserRole

	if err := c.ShouldBindJSON(&updateUserRole); err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	userRole, err := repo.UpdateUserRole(updateUserRole)

	if err != nil {
		ResponseData(c, http.StatusOK, nil)
		return
	}
	ResponseData(c, http.StatusOK, userRole)
}

func GetUserList(c *gin.Context) {
	userList, err := repo.GetUserList()

	if err != nil {
		ResponseData(c, http.StatusOK, nil)
		return
	}
	ResponseData(c, http.StatusOK, userList)
}

func UpdateUser(c *gin.Context) {

	var updateUser dto.UpdateUser
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	newUser, err := repo.UpdateUser(updateUser)

	if err != nil {
		ResponseData(c, http.StatusOK, err)
		return
	}
	ResponseData(c, http.StatusOK, newUser)
}

func DeleteUser(c *gin.Context) {
	id, _ := c.Params.Get("id")
	currentUser := c.MustGet("userId").(float64)

	if id == "" {
		ResponseError(c, http.StatusBadRequest, "Not exists")
		return
	}
	parseId, _ := strconv.ParseInt(id, 10, 0)

	if id == "1" || int64(currentUser) == parseId {
		ResponseError(c, http.StatusBadRequest, "Can not delete current user")
		return
	}

	ok, err := repo.DeleteUser(id)
	if err != nil {
		ResponseData(c, http.StatusBadRequest, "Can not delete user")
		return
	}
	ResponseData(c, http.StatusOK, ok)
}

func DumpUser(c *gin.Context) {
	err := repo.DumpUser()
	if err != nil {
		ResponseData(c, http.StatusOK, err)
		return
	}
	ResponseData(c, http.StatusOK, "Success")
}
