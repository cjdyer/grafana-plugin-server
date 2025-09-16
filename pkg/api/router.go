package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/repo.json", GetRepoJSON)

	// Plugin APIs
	r.GET("/api/plugins", GetPlugins)
	r.POST("/api/plugins", UploadPlugin)
	r.GET("/api/plugins/:slug", GetPluginBySlug)
	r.GET("/api/plugins/:slug/versions", GetVersions)
	r.GET("/api/plugins/:slug/versions/:ver", GetVersion)
	r.GET("/api/plugins/:slug/versions/:ver/download", DownloadVersion)
}
