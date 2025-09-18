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

	port := ":3838"
	log.Printf("Server starting on %s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
