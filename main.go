package main

import (
	"fmt"
	"os"

	env "github.com/joho/godotenv"
)

func main() {
	err := env.Load()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(os.Getenv("TEST"))
}
