package controllers

// Load libraries
import "github.com/gin-gonic/gin"

import "kurz/utils"

type Base int

// GET routes
func (b *Base) GETIndex(c *gin.Context) {
  c.HTML(200, "base.index.tmpl", map[string]interface{}{
    "title": "kurz",
  })
}

func (b *Base) GETAbout(c *gin.Context) {
  c.HTML(200, "base.about.tmpl", map[string]interface{}{
    "title": "kurz - about",
  })
}

func (b *Base) GETTerms(c *gin.Context) {
  c.HTML(200, "base.terms.tmpl", map[string]interface{}{
    "title": "kurz - terms",
  })
}

func (b *Base) GETLogin(c *gin.Context) {
  // Build render extras
  extra := map[string]interface{}{
    "title": "kurz - login",
  }

  // Check if session exists
  if c.GetString("session") != "" {
    // Load flash messages
    extra["error"] = utils.SessionFlash(c, "error")
    extra["info"]  = utils.SessionFlash(c, "info")
  }

  // Render page
  c.HTML(200, "base.login.tmpl", extra)
}

func (b *Base) GETRegister(c *gin.Context) {
  // Build render extras
  extra := map[string]interface{}{
    "title": "kurz - register",
  }

  // Check if session exists
  if c.GetString("session") != "" {
    // Load flash messages
    extra["error"] = utils.SessionFlash(c, "error")
    extra["info"]  = utils.SessionFlash(c, "info")
  }

  // Render page
  c.HTML(200, "base.register.tmpl", extra)
}

// POST routes
func (b *Base) POSTLogin(c *gin.Context) {
  // Get data
  inpEmail    := c.PostForm("email")
  inpPassword := c.PostForm("password")

  // Validate values
  if utils.ValidateEmail(inpEmail) && utils.ValidateLength(inpPassword, 8, 255) {
    // Get user from database
    rec := utils.GetUserByEmail(inpEmail, "password", "id")

    // Check for erro
    if rec == nil {
      // Set error message
      utils.SessionSet(c, "error", "wrong email")

      // Error response
      c.Redirect(302, "/login")
      return
    }

    // Get password
    password := rec.Bins["password"].(string)

    // Compare passwords
    compare := utils.BCryptCompare(password, inpPassword)

    // Check password
    if compare {
      // Update session
      utils.SessionSet(c, "user", rec.Bins["id"].(string))

      // Set succes message
      utils.SessionSet(c, "info", "welcome user")

      // Success response
      c.Redirect(302, "/board")
      return
    } else {
      // Set error message
      utils.SessionSet(c, "error", "wrong password")

      // Error response
      c.Redirect(302, "/login")
      return
    }

  } else {
    // Set error message
    utils.SessionSet(c, "error", "wrong email/password")

    // Error response
    c.Redirect(302, "/login")
    return
  }
}

func (b *Base) POSTRegister(c *gin.Context) {
  // Get data
  inpEmail    := c.PostForm("email")
  inpPassword := c.PostForm("password")

  // Validate values
  if utils.ValidateEmail(inpEmail) && utils.ValidateLength(inpPassword, 8, 255) {
    // Get user by email from database
    rec := utils.GetUserByEmail(inpEmail)

    // Check if account exists
    if rec != nil {
      // Set error message
      utils.SessionSet(c, "error", "email in use")

      // Error response
      c.Redirect(302, "/register")
      return
    }

    // Generate id
    id := utils.Identifier(inpEmail + c.ClientIP())

    // Encrypt password
    password := utils.BCrypt(inpPassword)

    // Build bins
    bins := utils.BinMap{
      "id": id,
      "email": inpEmail,
      "password": password,
    }

    // Build Database key
    key := utils.Key("users", id)

    // Insert into database
    err := utils.DB.Put(utils.WRITE, key, bins)

    // Check insert result
    if err != nil {
      // Set error message
      utils.SessionSet(c, "error", "try again later")

      // Error response
      c.Redirect(302, "/signup")
      return
    }

    // Set success message
    utils.SessionSet(c, "info", "account created")

    // Success response
    c.Redirect(302, "/login")
    return

  } else {
    // Set error message
    utils.SessionSet(c, "error", "form was not valid")

    // Error response
    c.Redirect(302, "/register")
    return
  }
}
