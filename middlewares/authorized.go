package middlewares

// Load libraries
import "github.com/gin-gonic/gin"

import "kurz/utils"

// Build the session handler function for gin
func Authorized() gin.HandlerFunc {
  // Build handler function
  return func(c *gin.Context) {
    // Get session field
    user := utils.SessionGet(c, "user")

    // Check value
    if user != "" {
      // Update status
      c.Set("user", user)
    } else {
      // Set status
      c.Redirect(302, "/login")

      // End the request
      c.Abort()
      return
    }

    // Jump to next function
    c.Next()
  }
}

func Guest() gin.HandlerFunc {
  // Build handler function
  return func(c *gin.Context) {
    // Get session field
    user := utils.SessionGet(c, "user")

    // Check value
    if user != "" {
      // Set status
      c.Redirect(302, "/u/board")

      // End the request
      c.Abort()
      return
    }

    // Jump to next function
    c.Next()
  }
}
