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
