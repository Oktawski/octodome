package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	authdom "octodome.com/api/internal/auth/domain"
	authprovider "octodome.com/api/internal/auth/provider"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JwtAuthMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userContext, err := extractUserFromJwt(r, db)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{"error": %q}`, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), authdom.UserContextKey, userContext)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func extractUserFromJwt(r *http.Request, db *gorm.DB) (*authdom.UserContext, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header required")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return nil, errors.New("authorization header format must be: Bearer {token}")
	}

	tokenStr := parts[1]

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&authdom.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*authdom.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	repo := authprovider.ProvideRoleReader(db)
	roles, err := repo.GetRolesByUserID(claims.UserID)
	if err != nil {
		return nil, err
	}

	return &authdom.UserContext{ID: claims.UserID, Roles: roles}, nil
}
