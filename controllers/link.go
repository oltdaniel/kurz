package controllers

// Load libraries
import "github.com/gin-gonic/gin"

type Link int

// GET routes
func (l *Link) GETLink(c *gin.Context) {
  c.String(200, "Hello World")
}

// POST routes
func (l *Link) POSTLink(c *gin.Context) {
  // Check if user is signed in
  if c.GetString("user") != "" {

  } else {

  }
}
