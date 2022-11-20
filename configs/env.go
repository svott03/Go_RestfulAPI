package configs

import (
	"fmt"
	"os"
)

func EnvMongoURI() string {
	// err := godotenv.Load()
  // if err != nil {
  //   log.Fatal("Error loading .env file")
  // }
	token := os.Getenv("MONGOURI")
	fmt.Println(token)
	return token
}