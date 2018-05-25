package utils

// Load standard libraries
import "encoding/base64"

// Load thrid party libraries
import as "github.com/aerospike/aerospike-client-go"
import "github.com/chilts/sid"

// Shared database client
var DB      = getClient()
var WRITE   = as.NewWritePolicy(0, 0)
var READ    = as.NewPolicy()
var SESSION = as.NewWritePolicy(0, 86400)
var QUERY   = as.NewQueryPolicy()

var WRITE_PUBLIC = as.NewWritePolicy(0, 1209600)
var WRITE_USER   = WRITE

type BinMap = as.BinMap
type DKey = as.Key

func getClient() *as.Client {
  // Create new client
  client, err := as.NewClient("127.0.0.1", 3000)

  // Check for error
  if err != nil {
    panic(err)
  }

  return client
}

func Key(set string, key string) *as.Key {
  askey, _ := as.NewKey("kurz", set, key)

  return askey
}

func Identifier(data string) string {
  return sid.Id() + base64.StdEncoding.EncodeToString([]byte(data))
}

func GetUserById(id string, fields ...string) *as.Record {
  // Build database statement
  stm := as.NewStatement("kurz", "users", "id")

  // Append fields
  stm.BinNames = append(stm.BinNames , fields...)

  // Add filter value to statement
  stm.Addfilter(as.NewEqualFilter("id", id))

  // Get records from database
  recordset, err := DB.Query(QUERY, stm)

  // Check for error
  if err != nil {
    return nil
  }

  // Get first record
  rec := <-recordset.Records

  // Return user record
  return rec
}

func GetUserByEmail(email string, fields ...string) *as.Record {
  // Build database statement
  stm := as.NewStatement("kurz", "users", "email")

  // Append fields
  stm.BinNames = append(stm.BinNames , fields...)

  // Add filter value to statement
  stm.Addfilter(as.NewEqualFilter("email", email))

  // Get records from database
  recordset, err := DB.Query(QUERY, stm)

  // Check for error
  if err != nil {
    return nil
  }

  // Get first record
  rec := <-recordset.Records

  // Return user record
  return rec
}

func GetLinks(author string) []map[string]interface{} {
  // Build database statement
  stm := as.NewStatement("kurz" ,"links", "author")

  // Append fields
  stm.BinNames = append(stm.BinNames, "slug", "visits")

  // Add filter value to statement
  stm.Addfilter(as.NewEqualFilter("author", author))

  // Get records from database
  recordset, err := DB.Query(QUERY, stm)

  // Check for error
  if err != nil {
    return nil
  }

  // Store the links in a new data structure
  links := make([]map[string]interface{}, 0)

  // Read all records
  for r := range recordset.Records {
    // Default value
    if r.Bins["visits"] == nil {
      r.Bins["visits"] = 0
    }

    // Build new data structure
    e := map[string]interface{}{
      "slug": r.Bins["slug"],
      "visits": r.Bins["visits"],
    }

    // Append new link format
    links = append(links, e)
  }

  // Return formatted links
  return links
}

func DeleteLink(author string, slug string) bool {
  // Build database key
  key := Key("links", slug)

  // Get record from database
  rec, err := DB.Get(READ, key)

  // Check for error
  if err != nil {
    return false
  }

  // Check if bin exists
  b, exists := rec.Bins["author"]

  if !exists {
    return false
  }

  // Check for author
  if b.(string) == author {
    // Delete from database
    existed, err := DB.Delete(WRITE, key)

    // Check for error
    if !existed || err != nil {
      return false
    }

    return true

  } else {
    return false
  }
}
