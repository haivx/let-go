package controller

import (
	"net/http"

	"final-project/dto"
	repo "final-project/repo"

	"github.com/gin-gonic/gin"
)

func CreatePermission(c *gin.Context) {
	var newPermission dto.CreatePermission

	if err := c.ShouldBindJSON(&newPermission); err != nil {
		ResponseError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	permission, err := repo.CreatePermission(newPermission)

	if err != nil {
		ResponseData(c, http.StatusOK, nil)
		return
	}
	ResponseData(c, http.StatusOK, permission)
}

func CreateRolePermission(c *gin.Context) {
	var createRolePermission dto.CreateRolePermission

	if err := c.ShouldBindJSON(&createRolePermission); err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	permission, err := repo.CreateRolePermission(createRolePermission)

	if err != nil {
		ResponseData(c, http.StatusOK, nil)
		return
	}

	ResponseData(c, http.StatusOK, permission)
}
