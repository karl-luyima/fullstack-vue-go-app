// Package middleware holds functions that run before your route handlers —
// things like "check this request has a valid token" that apply to many
// routes, so you write the check once here instead of repeating it.
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/utils"
)

// claimsKey is the key we use to stash the parsed token claims on the
// request, so later handlers in the same request can retrieve them.
const claimsKey = "authClaims"

// RequireAuth returns a Gin middleware function. Any route this is
// attached to will reject the request with 401 Unauthorized unless a
// valid "Authorization: Bearer <token>" header is present.
func RequireAuth(jwtManager *utils.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			c.Abort() // stops the request here — the real handler never runs
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		claims, err := jwtManager.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Stash the claims so downstream handlers (like UserController.Me)
		// can find out who's making this request.
		c.Set(claimsKey, claims)
		c.Next() // continue on to the actual route handler
	}
}

// MustGetClaims retrieves the claims set by RequireAuth. Only call this
// from a handler that's guaranteed to run after RequireAuth.
func MustGetClaims(c *gin.Context) *utils.Claims {
	val, _ := c.Get(claimsKey)
	return val.(*utils.Claims)
}