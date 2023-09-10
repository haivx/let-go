package middleware

import (
	repo "final-project/repo"
	services "final-project/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(action []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "status": http.StatusUnauthorized})
			return
		}

		claims, err := services.ValidateToken(strArr[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "status": http.StatusUnauthorized})
			return
		} else {
			email := claims["email"]
			listPermissions, _ := repo.GetUserPermission(email)
			listCode := []string{}
			for _, m := range *listPermissions {
				listCode = append(listCode, m.Code)
			}
			canAccess := hasItem(listCode, action)
			if !canAccess {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden", "status": http.StatusForbidden})
				return
			} else {
				c.Set("userId", claims["id"])
				c.Next()
			}
		}
	}
}

func hasItem(currentPermissions []string, actions []string) bool {
	for _, v := range currentPermissions {
		for _, u := range actions {
			if strings.EqualFold(v, u) {
				return true
			}
		}
	}
	return false
}
