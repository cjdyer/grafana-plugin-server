package api

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cjdyer/grafana-plugin-server/pkg/db"
	"github.com/cjdyer/grafana-plugin-server/pkg/plugins"
	"github.com/gin-gonic/gin"
)

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
		Type string `json:"type"`
		Name string `json:"name"`
		ID   string `json:"id"`
		Info struct {
			Description string `json:"description"`
			Author      struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"author"`
			Keywords []string `json:"keywords"`
			Version  string   `json:"version"`
		} `json:"info"`
	}

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	typeCode, err := GetTypeData(payload.Type)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plugin type"})
		return
	}

	t := time.Now().UTC()
	p := db.Plugin{
		Slug:        payload.ID,
		TypeId:      typeCode.Id,
		TypeName:    typeCode.Name,
		TypeCode:    typeCode.Code,
		Name:        payload.Name,
		URL:         payload.Info.Author.URL,
		Description: payload.Info.Description,
		OrgName:     payload.Info.Author.Name,
		OrgUrl:      payload.Info.Author.URL,
		Keywords:    payload.Info.Keywords,
		Version:     payload.Info.Version,
		UpdatedAt:   t.Format(time.RFC3339),
	}

	if err := plugins.AddPlugin(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Plugin uploaded successfully",
		"plugin":  p,
	})
}

func GetPluginBySlug(c *gin.Context) {
	slug := c.Param("slug")
	c.JSON(200, gin.H{"slug": slug})
}

func GetVersions(c *gin.Context) {
	slug := c.Param("slug")
	c.JSON(200, gin.H{"slug": slug, "versions": []string{}})
}

func GetVersion(c *gin.Context) {
	slug := c.Param("slug")
	ver := c.Param("ver")
	c.JSON(200, gin.H{"slug": slug, "version": ver})
}

func DownloadVersion(c *gin.Context) {
	slug := c.Param("slug")
	ver := c.Param("ver")
	c.JSON(200, gin.H{"message": "Download not implemented", "slug": slug, "version": ver})
}

type TypeMeta struct {
	Id   uint8
	Name string
	Code db.TypeCode
}

func GetTypeData(typeString string) (*TypeMeta, error) {
	switch typeString {
	case string(db.TypeCodeApp):
		return &TypeMeta{Id: 1, Name: "Application", Code: db.TypeCodeApp}, nil
	case string(db.TypeCodeDataSource):
		return &TypeMeta{Id: 2, Name: "Data Source", Code: db.TypeCodeDataSource}, nil
	case string(db.TypeCodePanel):
		return &TypeMeta{Id: 3, Name: "Panel", Code: db.TypeCodePanel}, nil
	}

	return nil, fmt.Errorf("datasource type cannot be empty")
}
