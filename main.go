package main

// Import libraries
import "github.com/gin-gonic/gin"

import "kurz/controllers"
import "kurz/middlewares"
import "kurz/utils"

import "os"

// Check if environment variable D is set
var DEBUG = (os.Getenv("D") == "1")

/**
* Package entry point
*/
func main() {
  // Set release mode when not in debug
  if !DEBUG { gin.SetMode(gin.ReleaseMode) }

  // Create new gin instance
  r := gin.New()

  // Add middlewares based on mode
  if DEBUG {
    // Logging
    r.Use(gin.Logger())

  } else {
    // Recovery
    r.Use(gin.Recovery())
  }

  // Build render engine
  render := utils.RenderEngine()

  // Compile templates
  render.Add("base.index.tmpl", "layout.main.tmpl")
  render.Add("base.about.tmpl", "layout.main.tmpl")
  render.Add("base.terms.tmpl", "layout.main.tmpl")
  render.Add("base.login.tmpl", "layout.main.tmpl")
  render.Add("base.register.tmpl", "layout.main.tmpl")

  render.Add("user.board.tmpl", "layout.user.tmpl")
  render.Add("user.link.tmpl", "layout.user.tmpl")

  // Assign render engine
  r.HTMLRender = render

  // Assign session middleware
  r.Use(middlewares.Session())

  // Create controllers
  baseController := new(controllers.Base)
  userController := new(controllers.User)
  linkController := new(controllers.Link)

  // Group guest only routes
  guestGroup := r.Group("/")

  // Add guest only middleare to guest routes
  guestGroup.Use(middlewares.Guest())

  // Add routes
  {
    guestGroup.GET("/", baseController.GETIndex)
    guestGroup.GET("/about", baseController.GETAbout)
    guestGroup.GET("/terms", baseController.GETTerms)
    guestGroup.GET("/login", baseController.GETLogin)
    guestGroup.GET("/register", baseController.GETRegister)

    guestGroup.POST("/login", baseController.POSTLogin)
    guestGroup.POST("/register", baseController.POSTRegister)
  }

  // Group authorized only routes
  authorizedGroup := r.Group("/")

  // Add authorized only middleware to authorized group
  authorizedGroup.Use(middlewares.Authorized())

  // Add routes
  {
    authorizedGroup.GET("/board", userController.GETBoard)
    authorizedGroup.GET("/link/:short", userController.GETLink)
  }

  // Handle calls
  r.POST("/links", linkController.POSTLink)
  r.GET("/l/:short", linkController.GETLink)

  // Start instance
  r.Run(":4001")
}
