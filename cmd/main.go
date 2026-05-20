package main

import (
	"fmt"
	"net/http"

	"unicore/config"
)

func main() {
	config.DbConnection()
	db := config.DB

	fmt.Println("we got a server")
}
