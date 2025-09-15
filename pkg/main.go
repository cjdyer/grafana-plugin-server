package main

import (
	"log"

	"github.com/cjdyer/grafana-plugin-server/pkg/api"
	"github.com/cjdyer/grafana-plugin-server/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	r := gin.Default()
	api.RegisterRoutes(r)

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
