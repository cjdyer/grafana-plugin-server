package api

import (
	"io"
	"net/http"

	"github.com/cjdyer/grafana-plugin-server/pkg/db"
	"github.com/cjdyer/grafana-plugin-server/pkg/plugins"
	"github.com/gin-gonic/gin"
)

func GetRepoJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"plugins": []string{},
	})
}

func ProxyGrafanaAPI(c *gin.Context) {
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
}

func GetPlugins(c *gin.Context) {
	plugins, err := plugins.ListPlugins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": plugins,
	})
}

func UploadPlugin(c *gin.Context) {
	var payload struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		URL  string `json:"url"`
		Name string `json:"name"`
	}

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	p := db.Plugin{ID: payload.ID, Type: db.Type(payload.Type), Name: payload.Name, URL: payload.URL}

	if err := plugins.AddPlugin(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Plugin uploaded successfully",
		"plugin":  p,
	})
}
