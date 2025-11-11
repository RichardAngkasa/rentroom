package middleware

import (
	"context"
	"fmt"
	"net/http"
	"rentroom/utils"
)

func JwtAuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.ExtractTokenFromHeader(r)
		if err != nil {
			utils.JSONError(w, "unauthorized "+err.Error(), http.StatusUnauthorized)
			return
		}
		claims, err := ValidateToken(token, "user")
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		userID := int(claims["id"].(float64))
		redisKey := fmt.Sprintf("session:user:%d", userID)

		storedToken, err := utils.RedisUser.Get(utils.Ctx, redisKey).Result()
		if err != nil || storedToken != token {
			utils.JSONError(w, "session expired or invalid", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), CtxUserID, userID)
		ctx = context.WithValue(ctx, CtxRole, "user")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func JwtAuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.ExtractTokenFromHeader(r)
		if err != nil {
			utils.JSONError(w, "unauthorized "+err.Error(), http.StatusUnauthorized)
			return
		}
		claims, err := ValidateToken(token, "admin")
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		adminID := int(claims["id"].(float64))
		redisKey := fmt.Sprintf("session:admin:%d", adminID)

		storedToken, err := utils.RedisUser.Get(utils.Ctx, redisKey).Result()
		if err != nil || storedToken != token {
			utils.JSONError(w, "session expired or invalid", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), CtxAdminID, adminID)
		ctx = context.WithValue(ctx, CtxRole, "admin")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
