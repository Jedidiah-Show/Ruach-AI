package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"ruach-ai/database"
	"ruach-ai/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	routes.RegisterHealthRoutes(router)

	conn, err := database.Connect()

	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	fmt.Printf("Starting Ruach AI backend on port %s\n", port)

	router.Run(":" + port)
}
