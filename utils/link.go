package utils

import as "github.com/aerospike/aerospike-client-go"
import nanoid "github.com/matoous/go-nanoid"

const ALPHABET = "abcdefghijklmnopqrstuvwxyz0123456789#._-:!"

func LinkShort() (string, *as.Key) {
  // Generate random nano id
  id, err := nanoid.Generate(ALPHABET, 6)

  // Check for error
  if err != nil {
    // Generate new id
    return LinkShort()
  }

  // Build database key
  key := Key("links", id)

  // Check if id exists
  exists, err := DB.Exists(READ, key)

  // Check for error
  if exists || err != nil {
    // Generate new id
    return LinkShort()
  }

  return id, key
}
