// docs https://goethereumbook.org/transfer-eth/

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	private_key_account_0 := os.Getenv("PRIVATE_KEY_ACCOUNT_0")
	private_key_account_1 := os.Getenv("PRIVATE_KEY_ACCOUNT_1")

	fmt.Println("private_key_account_0: ", private_key_account_0)
	fmt.Println("private_key_account_1: ", private_key_account_1)
}
