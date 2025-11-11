package admin

import (
	"fmt"
	"net/http"
	"os"
	"rentroom/internal/models"
	"rentroom/utils"
	"time"

	"gorm.io/gorm"
)

func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		var req models.AdminLoginRequest
		err := utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.FieldChecker(req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		adminUsername := os.Getenv("ADMIN_USERNAME")
		if req.Username != adminUsername {
			utils.JSONError(w, "unauthorized", http.StatusBadRequest)
			return
		}
		adminPassword := os.Getenv("ADMIN_PASSWORD")
		if req.Password != adminPassword {
			utils.JSONError(w, "unauthorized", http.StatusBadRequest)
			return
		}

		// QUERY
		token, err := utils.GenerateJWT(uint(1), "admin")
		if err != nil {
			utils.JSONError(w, "token generation failed", http.StatusInternalServerError)
			return
		}
		err = utils.RedisUser.Set(utils.Ctx,
			"session:admin:"+fmt.Sprint(1),
			token,
			24*time.Hour,
		).Err()
		if err != nil {
			utils.JSONError(w, "redis", http.StatusInternalServerError)
			return
		}
		utils.SetTokenInHeader(w, token)

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "admin logged in",
			Data: models.AdminLoginResponse{
				Token: token,
			},
		}, http.StatusCreated)
	}
}
