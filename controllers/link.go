package controllers

// Load libraries
import "github.com/gin-gonic/gin"

type Link int

// GET routes
func (l *Link) GETLink(c *gin.Context) {
  c.String(200, "Hello World")
}
