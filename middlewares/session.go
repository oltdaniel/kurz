package middlewares

// Load thir party libraries
import "github.com/gin-gonic/gin"

// Build the session handler function for gin
func Session() gin.HandlerFunc {
  // Build handler function
  return func(c *gin.Context) {
    // Check for session id in cookie
    cookie, err := c.Cookie("session")

    // Check for error
    if err == nil {
      // Set session id in request
      c.Set("session", cookie)
    }

    // Jump to next function
    c.Next()
  }
}
