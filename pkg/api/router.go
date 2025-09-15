package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/repo.json", GetRepoJSON)
	r.GET("/api/plugins", ProxyGrafanaAPI)
    
	r.POST("/api/plugins", UploadPlugin)
}
