package main

import (
	"fmt"
	"log"

	"github.com/ouhabmoh/HR/initializers"
	"github.com/ouhabmoh/HR/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {

	initializers.DB.AutoMigrate(&models.User{}, &models.Job{}, &models.Application{})
	fmt.Println("👍 Migration complete")
}
