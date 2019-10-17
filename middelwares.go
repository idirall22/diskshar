package userAccount

import (
	"context"
	"net/http"
)

// Key used in context
type Key string

// IDCtx id user stored in context
var IDCtx Key = "user_id"

// UsernameCtx username user stored in context
var UsernameCtx Key = "user_username"

// AuthnticateUser check if user is authonticated resquest
func AuthnticateUser(fn http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		id, username, err := validateBearerToken(tokenString)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		ctx := storeUserInContext(r.Context(), id, username)
		r = r.WithContext(ctx)

		fn.ServeHTTP(w, r)
	})
}

// storeUserInContext set id an username to requets context
func storeUserInContext(ctx context.Context, id int64, username string) context.Context {

	ctx = context.WithValue(ctx, IDCtx, id)
	ctx = context.WithValue(ctx, UsernameCtx, username)
	return ctx
}
