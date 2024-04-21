package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/maskedeman/hackernews/internal/users"
	"github.com/maskedeman/hackernews/pkg/jwt"
)

var usrCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			user := users.User{Username: username}
			id, err := users.GetUserIDByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			user.ID = strconv.Itoa(id)
			ctx := context.WithValue(r.Context(), usrCtxKey, &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(usrCtxKey).(*users.User)
	return raw
}
