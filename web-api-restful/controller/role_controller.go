package controller

import (
	dto "final-project/dto"
	repo "final-project/repo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateRole(c *gin.Context) {
	var newRole dto.CreateRole
	if err := c.ShouldBindJSON(&newRole); err != nil {
		ResponseError(c, http.StatusBadRequest, "Invalid input")
		return
	}
	v := validator.New()

	err := v.Struct(newRole)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		if validationErrors != nil {
			ResponseError(c, http.StatusBadRequest, "Role is one of ADMIN, USER, or MOD")
			return
		}
	}

	role, err := repo.CreateRole(&newRole)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "Error insert to database")
		return
	}

	ResponseData(c, http.StatusOK, role)
}
