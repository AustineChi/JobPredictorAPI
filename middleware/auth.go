package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


// validates the token, extracts UserID and sets it in the request context
func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		//validate token and exract userID
		userID, err := ValidateToken(authHeader)
		if err != nil {
			log.Println("invalid or expired token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		// Set the user ID in the request context
		c.Set("userID", userID)

		// Continue processing the request
		c.Next()
	}
}