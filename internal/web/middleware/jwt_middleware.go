// package middleware

// import (
// 	"net/http"
// 	auth "octodome/internal/auth/domain"
// 	"os"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// func JwtAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
// 			return
// 		}

// 		parts := strings.SplitN(authHeader, " ", 2)
// 		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be: Bearer {token}"})
// 			return
// 		}

// 		tokenStr := parts[1]

// 		token, err := jwt.ParseWithClaims(
// 			tokenStr,
// 			&auth.UserClaims{},
// 			func(token *jwt.Token) (interface{}, error) {
// 				return jwtSecret, nil
// 			})
// 		if err != nil || !token.Valid {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
// 			return
// 		}

// 		claims, ok := token.Claims.(*auth.UserClaims)
// 		if !ok {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
// 			return
// 		}

// 		userContext := &auth.UserContext{ID: claims.UserID}

// 		c.Set("userContext", userContext)

// 		c.Next()
// 	}
// }

package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	auth "octodome/internal/auth/domain"
	authdom "octodome/internal/auth/domain"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userContext, err := extractUserFromJwt(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error: %q}`, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), authdom.UserContextKey, userContext)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractUserFromJwt(r *http.Request) (*auth.UserContext, error) {
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
		&auth.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*auth.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return &auth.UserContext{ID: claims.UserID}, nil
}
