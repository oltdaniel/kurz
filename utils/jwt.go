package utils

// Load libraries
import "github.com/dgrijalva/jwt-go"

import "io/ioutil"
import "fmt"
import "time"

var SECRET = ReadKey()

func ReadKey() string {
  // Read key file
  dat, err := ioutil.ReadFile("key")

  // Throw error
  if err != nil {
    panic(err)
    return ""
  }

  // Return read data
  return string(dat)
}

func JWTParse(tokenString string) (jwt.MapClaims, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Undefined signing method: %v", token.Header["alg"])
    }

    return SECRET, nil
  })

  if err != nil {
    return nil, fmt.Errorf("Invalid token")
  }

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    return claims, nil
  } else {
    return nil, fmt.Errorf("Invalid token")
  }
}

func JWTBuild(user string) string {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "user": user,
    "time": time.Now().Unix(),
  })

  tokenString, err := token.SignedString([]byte(SECRET))

  if err != nil {
    panic(err)
  }

  return tokenString
}
