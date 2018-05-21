package utils

// Load third party libraries
import "golang.org/x/crypto/bcrypt"

// Generate bcrypt hash from password
func BCrypt(password string) string {
  hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

  if err != nil {
    panic(err)
  }

  return string(hash)
}

// Compate an hash with a password
func BCryptCompare(hash string, password string) bool {
  return (bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil)
}
