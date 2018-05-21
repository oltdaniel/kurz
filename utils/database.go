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
