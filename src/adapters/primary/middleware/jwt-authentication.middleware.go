package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	hexagon_repository "pitch-hunt/src/hexagon/gateways/repository"

	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationMiddleware struct {
	userRepository hexagon_repository.UserRepository
}

func NewAuthenticationMiddleware(userRepository hexagon_repository.UserRepository) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		userRepository: userRepository,
	}
}

func (am *AuthenticationMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user, err := am.userRepository.FindById(claims["id"].(string))
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
			}

			ctx := context.WithValue(r.Context(), "user", user)
			r = r.WithContext(ctx)
		} else {
			fmt.Println("Invalid token or claims")
		}

		next.ServeHTTP(w, r)
	})
}
