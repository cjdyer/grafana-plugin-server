package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Version struct {
	Version string `json:"version"`
	URL     string `json:"url"`
}

type Plugin struct {
	ID       string    `json:"id"`
	Type     string    `json:"type"`
	Versions []Version `json:"versions"`
}

func main() {
	r := gin.Default()

	r.GET("/repo.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"plugins": []Plugin{
				{
					ID:   "plugin",
					Type: "panel",
					Versions: []Version{
						{Version: "1.0.0", URL: "http://localhost:8080/plugins/plugin"},
					},
				},
			},
		})
	})

	r.GET("/api/plugins", func(c *gin.Context) {
		resp, err := http.Get("https://grafana.com/api/plugins")
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to reach Grafana API"})
			return
		}
		defer resp.Body.Close()

		c.Status(resp.StatusCode)
		c.Header("Content-Type", resp.Header.Get("Content-Type"))
		body, _ := io.ReadAll(resp.Body)
		c.Writer.Write(body)
	})

	r.Run(":8080")
}
