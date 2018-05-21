package controllers

// Load libraries
import "github.com/gin-gonic/gin"

type Base int

// GET routes
func (b *Base) GETIndex(c *gin.Context) {
  c.HTML(200, "base.index.tmpl", map[string]interface{}{
    "title": "kurz",
  })
}

func (b *Base) GETAbout(c *gin.Context) {
  c.String(200, "Hello World")
}

func (b *Base) GETLogin(c *gin.Context) {
  c.String(200, "Hello World")
}

func (b *Base) GETRegister(c *gin.Context) {
  c.String(200, "Hello World")
}

// POST routes
func (b *Base) POSTLogin(c *gin.Context) {
  c.String(200, "Hello World")
}

func (b *Base) POSTRegister(c *gin.Context) {
  c.String(200, "Hello World")
}
