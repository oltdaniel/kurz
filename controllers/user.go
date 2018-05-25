package controllers

// Load libraries
import "github.com/gin-gonic/gin"

import "kurz/utils"

import "time"

type User int

// GET routes
func (u *User) GETBoard(c *gin.Context) {
  // Get token
  token := c.GetString("token")

  // Check if token exists
  if token == "" {
    // Build new token
    token = utils.JWTBuild(c.GetString("user"))

    // Assign token
    c.Set("token", token)
  }

  // Get links
  links := utils.GetLinks(c.GetString("user"))

  c.HTML(200, "user.board.tmpl", map[string]interface{}{
    "title": "kurz",
    "token": token,
    "links": links,
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
func (u *User) POSTApiLink(c *gin.Context) {
  // Get data
  inpLink   := c.PostForm("link")
  inpSlug   := c.PostForm("slug")
  inpToken  := c.Query("token")

  // Validate jwt
  token, err := utils.JWTParse(inpToken)

  // Check for error
  if err != nil {
    // Error response
    c.JSON(403, map[string]interface{}{
      "error": true,
      "message": "invalid token",
    })
    return
  }

  // Validate url
  if !utils.ValidateLink(inpLink) {
    // Error response
    c.JSON(400, map[string]interface{}{
      "error": true,
      "message": "invalid link",
    })
    return

  } else {
    // Rmemeber key properties
    var key *utils.DKey
    var random string

    // Check for slug
    if inpSlug != "" {
      // Assing random
      random = inpSlug

      // Build database key
      key = utils.Key("links", inpSlug)

      // Check if key exists
      exists, err := utils.DB.Exists(utils.READ, key)

      // Check for error
      if exists || err != nil {
        // Error response
        c.JSON(400, map[string]interface{}{
          "error": true,
          "message": "slug in use",
        })
        return
      }

    } else {
      // Generate random link hash
      random, key = utils.LinkShort()
    }

    // Buld bins
    bins := utils.BinMap{
      "link": inpLink,
      "slug": random,
      "created_at": time.Now().Unix(),
      "author": token["user"],
      "ip": c.ClientIP(),
    }

    // Insert into database
    err := utils.DB.Put(utils.WRITE_USER, key, bins)

    // Check insert result
    if err != nil {
      // Error response
      c.JSON(500, map[string]interface{}{
        "error": true,
        "message": "try again later",
      })
      return
    }

    // Success response
    c.JSON(200, map[string]interface{}{
      "error": false,
      "message": "link inserted",
      "data": random,
    })
    return
  }
}

// DELETE routes
func (u *User) DELETEApiLink(c *gin.Context) {
  // Get data
  inpSlug  := c.Param("short")
  inpToken := c.Query("token")

  // Validate jwt
  token, err := utils.JWTParse(inpToken)

  // Check for error
  if err != nil {
    // Error response
    c.JSON(403, map[string]interface{}{
      "error": true,
      "message": "Invalid token",
    })
    return
  }

  // Delete link
  success := utils.DeleteLink(token["user"].(string), inpSlug)

  // Check for success
  if success {
    // Success response
    c.JSON(200, map[string]interface{}{
      "error": false,
      "message": "link deleted",
    })
    return
  }

  // Error response
  c.JSON(400, map[string]interface{}{
    "error": true,
    "message": "link does not exist",
  })
  return
}
