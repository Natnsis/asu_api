package main

import (
	"fmt"

	"unicore/config"
)

func main() {
	config.DbConnection()
	db := config.DB

	fmt.Println("we got a server")
}
