package utils

// Load third party libraries
import "github.com/gin-gonic/gin"
import as "github.com/aerospike/aerospike-client-go"

// Check for session id
func SessionCheck(c *gin.Context) {
  if c.GetString("session") == "" {
    // Generate session id
    sessionid := Identifier(c.ClientIP())

    // Update session id data
    c.SetCookie("session", sessionid, 86400, "/", "", false, false)
    c.Set("session", sessionid)
  }
}

// Set value based on session
func SessionSet(c *gin.Context, name string, value string) {
  // Check for session id
  SessionCheck(c)

  // Get session id
  sessionid := c.GetString("session")

  // Build session key
  key := Key("sessions", sessionid)

  // Construct bins
  bins := as.BinMap{
    name: value,
  }

  // Set value
  DB.Put(SESSION, key, bins)
}

// Get value based on session
func SessionGet(c *gin.Context, name string) string {
  // Check for session id
  SessionCheck(c)

  // Get session id
  sessionid := c.GetString("session")

  // Build session key
  key := Key("sessions", sessionid)

  // Set value
  rec, err := DB.Get(READ, key)

  // Check for error
  if err == nil && rec != nil && rec.Bins[name] != nil {
    // Return found bin
    return rec.Bins[name].(string)
  }

  // Return empty on error
  return ""
}

// Get and remove value based on session
func SessionFlash(c *gin.Context, name string) string {
  // Get value
  r := SessionGet(c, name)

  // Clear property
  SessionSet(c, name, "")

  // Return flash value
  return r
}

func SessionDelete(c *gin.Context) {
  // Check for session id
  SessionCheck(c);

  // Build key
  key := Key("sessions", c.GetString("session"))

  // Delete session from database
  DB.Delete(WRITE, key)

  // Remove cookie and context string
  c.SetCookie("session", "", -1, "/", "", false, false)
  c.Set("session", "")
}
