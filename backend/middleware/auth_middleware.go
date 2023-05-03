package middleware

import (
	"context"
	"golang-nextjs-app/gateway/firebase"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		if idToken == "" {
			http.Error(w, "Bearer token is required", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		token, err := firebase.VerifyIDToken(ctx, idToken)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, "uid", token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
