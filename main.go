package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/narozhniyy/test/resources"
	"github.com/narozhniyy/test/router"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Initialize db connection
	resources.ConnectDB()
	defer func() {
		if err := resources.CloseDB(); err != nil {
			log.Fatal(err)
		}
	}()
	// Initialize router with echo framework
	e := router.InitRouter()
	// Start server
	e.Logger.Fatal(e.Start(":8082"))
}
