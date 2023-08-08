package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"suggester/api/routers"
	"suggester/db"
)

func init() {
	err := db.InitSchema()
	if err != nil {
		fmt.Println("Cant create schema")
	}
}

func main() {
	routers.InitRouters()
}
