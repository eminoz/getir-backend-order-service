package main

import (
	"fmt"

	"github.com/eminoz/getir-backend-order-service/pkg/config"
	"github.com/eminoz/getir-backend-order-service/pkg/database"
	"github.com/eminoz/getir-backend-order-service/pkg/router"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		fmt.Printf("config.Setup() error: %s", err)
	}
	database.Setup()

	r := router.Setup()
	r.Listen(":3034")
}
