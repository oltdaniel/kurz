package controllers

// Load libraries
import "github.com/gin-gonic/gin"

import "kurz/utils"

import "time"

type Link int

// GET routes
func (l *Link) GETLink(c *gin.Context) {
  // Read url
  paraShort := c.Param("short")

  // Build database key
  key := utils.Key("links", paraShort)

  // Get link from database
  rec, err := utils.DB.Get(utils.READ, key)

  // Check for error
  if rec == nil || err != nil {
    // Error response
    c.Redirect(302, "/")
    return
  }

  // Build bins
  bins := utils.BinMap{
    "visits": 1,
  }

  // Increment visit count
  go utils.DB.Add(utils.WRITE, key, bins)

  // Redirect
  c.Redirect(302, rec.Bins["link"].(string))
  return
}

// POST routes
func (l *Link) POSTApiLink(c *gin.Context) {
  // Get data
  inpLink := c.PostForm("link")

  // Validate link
  if !utils.ValidateLink(inpLink) {
    // Error response
    c.JSON(400, map[string]interface{}{
      "error": true,
      "message": "invalid link",
    })
    return

  } else {
    // Generate random link hash
    random, key := utils.LinkShort()

    // Build bins
    bins := utils.BinMap{
      "link": inpLink,
      "slug": random,
      "created_at": time.Now().Unix(),
      "ip": c.ClientIP(),
    }

    // Insert into database
    err := utils.DB.Put(utils.WRITE_PUBLIC, key, bins)

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
