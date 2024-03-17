package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ValidationMiddleware struct {
	validate *validator.Validate
}

func NewValidationMiddleware() *ValidationMiddleware {
	return &ValidationMiddleware{
		validate: validator.New(),
	}
}

func (mw *ValidationMiddleware) Handle(dto interface{}, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dtoValue := reflect.New(reflect.TypeOf(dto)).Interface()

		err := json.NewDecoder(r.Body).Decode(dtoValue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = mw.validate.Struct(dtoValue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "dto", dtoValue)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
