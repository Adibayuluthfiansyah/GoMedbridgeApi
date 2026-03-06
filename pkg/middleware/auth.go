package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Adibayuluthfiansyah/GoMedbridgeApi/pkg/response"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"
const UserRoleKey contextKey = "user_role"

func Auth(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
					Status:  "error",
					Message: "Missing Authorized Header",
				})
				return
			}
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
					Status:  "error",
					Message: "Invalid Authorized Header",
				})
				return
			}
			tokenString := parts[1]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, http.ErrAbortHandler
				}
				return []byte(jwtSecret), nil
			})
			if err != nil || !token.Valid {
				response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
					Status:  "error",
					Message: "Invalid or expired token",
				})
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				response.WriteJSON(w, http.StatusUnauthorized, response.JSONResponse{
					Status:  "error",
					Message: "Invalid token claim",
				})
				return
			}
			userID := claims["userID"].(string)
			userRole := claims["role"].(string)

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			ctx = context.WithValue(ctx, UserRoleKey, userRole)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
