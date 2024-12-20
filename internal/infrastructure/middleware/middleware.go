package middleware

import (
	"context"
	"loan-service/internal/entity"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type contextKey string

const UserIDContextKey contextKey = "userID"
const UserRoleContextKey contextKey = "userRole"

var skippedPaths = []string{"/users"}

// JWTMiddlewareWithDB demonstrating a middleware that validates the token, fetches user details from the database, and adds it to the context.
func JWTMiddlewareWithDB(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, path := range skippedPaths {
				if strings.HasPrefix(r.URL.Path, path) {
					next.ServeHTTP(w, r)
					return
				}
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
				return
			}

			userID, err := strconv.ParseInt(parts[1], 10, 64) // Assuming the user ID is in the token
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			user, err := fetchUserFromDB(db, userID)
			if err != nil {
				log.Println("Error fetching user:", err)
				http.Error(w, "user not found", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDContextKey, userID)
			ctx = context.WithValue(ctx, UserRoleContextKey, user.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func fetchUserFromDB(db *gorm.DB, userID int64) (*entity.User, error) {
	var user entity.User
	err := db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
