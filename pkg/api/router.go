package api

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/plugins", GetPlugins)
		api.POST("/plugins", UploadPlugin)
		api.GET("/plugins/:slug", GetPluginBySlug)
		api.GET("/plugins/:slug/versions", GetVersions)
		api.GET("/plugins/:slug/versions/:ver", GetVersion)
		api.GET("/plugins/:slug/versions/:ver/logos/:variant", GetLogo)
		api.GET("/plugins/:slug/versions/:ver/download", DownloadVersion)

		api.GET("/plugins/versioncheck", VersionCheck)
	}

	r.NoRoute(func(c *gin.Context) {
		reqPath := c.Request.URL.Path
		fullPath := filepath.Join("dist", reqPath)

		if fi, err := os.Stat(fullPath); err == nil && !fi.IsDir() {
			c.File(fullPath)
			return
		}

		c.File("dist/index.html")
	})
}
