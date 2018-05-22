package controllers

// Load libraries
import "github.com/gin-gonic/gin"

import "kurz/utils"

type Link int

// GET routes
func (l *Link) GETLink(c *gin.Context) {
  c.String(200, "Hello World")
}

// POST routes
func (l *Link) POSTLink(c *gin.Context) {
  // Get data
  inpLink := c.PostForm("link")

  // Generate random link hash
  random := utils.LinkShort(inpLink)

  c.String(200, random)
}
