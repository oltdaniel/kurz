package controllers

// Load libraries
import "github.com/gin-gonic/gin"

import "kurz/utils"

type User int

// GET routes
func (u *User) GETBoard(c *gin.Context) {
  c.HTML(200, "user.board.tmpl", map[string]interface{}{
    "title": "kurz",
  })
}

func (u *User) GETLink(c *gin.Context) {
  c.HTML(200, "user.link.tmpl", map[string]interface{}{
    "title": "kurz - link",
  })
  return
}

func (u *User) GETLogout(c *gin.Context) {
  // Delete session
  utils.SessionDelete(c)

  // Set success message
  utils.SessionSet(c, "info", "signed out")

  // Success response
  c.Redirect(302, "/login")
  return
}

// POST routes
func (u *User) POSTLink(c *gin.Context) {
  c.String(200, "Hello World")
  return
}
