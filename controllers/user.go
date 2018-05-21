package controllers

// Load libraries
import "github.com/gin-gonic/gin"

type User int

// GET routes
func (u *User) GETBoard(c *gin.Context) {
  c.String(200, "Hello World")
}

func (u *User) GETLink(c *gin.Context) {
  c.String(200, "Hello World")
}

// POST routes
func (u *User) POSTLink(c *gin.Context) {
  c.String(200, "Hello World")
}
