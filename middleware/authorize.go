package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ouhabmoh/HR/models"
)

type AllowedRoles struct {
	Roles []string
}

func NewAllowedRoles(roles ...string) *AllowedRoles {
	return &AllowedRoles{Roles: roles}
}

const (
	RoleCandidate = "candidate"
	RoleRecruiter = "recruiter"
)

func AuthorizationMiddleware(allowedRoles *AllowedRoles) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user's role from the JWT token or session
		userRole := c.MustGet("currentUser").(models.User).Role

		// Check if the user's role is among the allowed roles
		isAllowed := false
		for _, role := range allowedRoles.Roles {
			if role == userRole {
				isAllowed = true
				break
			}
		}

		// If the user is not allowed, return an unauthorized error
		if !isAllowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are Not Allowed to perform this action"})
			return
		}

		// If the user is allowed, continue to the next handler
		c.Next()
	}
}

func AuthorizeRoles(roles ...string) gin.HandlerFunc {
	return AuthorizationMiddleware(NewAllowedRoles(roles...))
}
