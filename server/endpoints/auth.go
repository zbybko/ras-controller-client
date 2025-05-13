package endpoints

import (
	"net/http"
	"ras/storage"
	"ras/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Password string `json:"password"`
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		storedPasswordHash := storage.GetPassword()

		if err := bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(request.Password)); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		token, err := utils.GenerateJWT()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"token": token, "success": true})
	}
}
func ChangePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			OldPassword string `json:"oldPassword"`
			NewPassword string `json:"newPassword"`
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		storedPasswordHash := storage.GetPassword()

		if err := bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(request.OldPassword)); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate password hash"})
			return
		}

		storage.SetPassword(string(newPasswordHash))

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
