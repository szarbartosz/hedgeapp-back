package validators

import (
	"main/initializers"
	"main/models"

	"github.com/gin-gonic/gin"
)

func UserAlreadyExists(c *gin.Context, body models.Credentials) bool {
	var user models.User
	result := initializers.DB.First(&user, "email = ?", body.Email)

	return result.RowsAffected != 0
}
