package middleware

import (
	"net/http"
	"strings"

	"github.com/Uranury/WorkoutApp/internal/auth"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type contextKey string

var claimsKey = contextKey("claimsKey")

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), claimsKey, claims)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func GetUserID(ctx context.Context) (uuid.UUID, bool) {
	claims, ok := ctx.Value(claimsKey).(auth.Claims)
	if !ok {
		return uuid.UUID{}, false
	}

	userIDRaw, ok := claims["user_id"]
	if !ok {
		return uuid.UUID{}, false
	}

	userIDStr, ok := userIDRaw.(string)
	if !ok {
		return uuid.UUID{}, false
	}

	id, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.UUID{}, false
	}

	return id, true
}

func GetUserRole(ctx context.Context) (auth.Role, bool) {
	claims, ok := ctx.Value(claimsKey).(auth.Claims)
	if !ok {
		return "", false
	}

	roleRaw, ok := claims["role"]
	if !ok {
		return "", false
	}

	roleStr, ok := roleRaw.(string)
	if !ok {
		return "", false
	}

	return auth.Role(roleStr), true
}
