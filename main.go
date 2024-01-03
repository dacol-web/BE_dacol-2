package main

import (
	"os"

	"github.com/Hy-Iam-Noval/dacol-2/src"
	_ "github.com/joho/godotenv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := src.Route()

	r.Listen(":" + port)
}
