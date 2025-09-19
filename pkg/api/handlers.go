package api

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cjdyer/grafana-plugin-server/pkg/db"
	"github.com/cjdyer/grafana-plugin-server/pkg/plugins"
	"github.com/gin-gonic/gin"
)

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
	file, err := c.FormFile("plugin")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Plugin file is required"})
		return
	}

	if !strings.HasSuffix(file.Filename, ".tar") && !strings.HasSuffix(file.Filename, ".tar.gz") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only .tar or .tar.gz files are allowed"})
		return
	}

	tempPath := "/tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	metadata, err := ExtractPluginMetadata(tempPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plugin package: " + err.Error()})
		return
	}

	readme, err := ExtractPluginReadme(tempPath)
	if err != nil {
		log.Println("Plugin Readme invalid", err)
	}

	typeCode, err := GetTypeData(metadata.Type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plugin type"})
		return
	}

	t := time.Now().UTC()
	p := db.Plugin{
		Slug:        metadata.ID,
		TypeId:      typeCode.Id,
		TypeName:    typeCode.Name,
		TypeCode:    typeCode.Code,
		Name:        metadata.Name,
		URL:         metadata.Info.Author.URL,
		Description: metadata.Info.Description,
		OrgName:     metadata.Info.Author.Name,
		OrgUrl:      metadata.Info.Author.URL,
		Keywords:    metadata.Info.Keywords,
		Version:     metadata.Info.Version,
		UpdatedAt:   t.Format(time.RFC3339),
		Readme:      readme,
		FilePath:    "/plugins/" + file.Filename,
	}

	existing, err := plugins.GetPluginBySlug(metadata.ID)
	if err == nil && existing != nil {
		if err := plugins.UpdatePlugin(p); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := plugins.AddPlugin(p); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	pluginDir := "./static/plugins/" + metadata.ID
	if err := os.MkdirAll(pluginDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create plugin dir"})
		return
	}

	SaveLogos(tempPath, metadata, pluginDir)
	os.Rename(tempPath, pluginDir+"/dist.tar")

	c.JSON(http.StatusCreated, gin.H{
		"message": "Plugin uploaded successfully",
		"plugin":  p,
	})
}

func GetPluginBySlug(c *gin.Context) {
	slug := c.Param("slug")
	plugin, err := plugins.GetPluginBySlug(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, plugin)
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

func ExtractPluginMetadata(tarPath string) (*db.Payload, error) {
	f, err := os.Open(tarPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var fileReader io.Reader = f
	if strings.HasSuffix(tarPath, ".gz") {
		if fileReader, err = gzip.NewReader(f); err != nil {
			return nil, err
		}
	}

	tarReader := tar.NewReader(fileReader)
	var metadata db.Payload

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if strings.HasSuffix(header.Name, "plugin.json") {
			data, err := io.ReadAll(tarReader)
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(data, &metadata); err != nil {
				return nil, err
			}
			return &metadata, nil
		}
	}

	return nil, errors.New("plugin.json not found in archive")
}

func ExtractPluginReadme(tarPath string) (string, error) {
	f, err := os.Open(tarPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var fileReader io.Reader = f
	if strings.HasSuffix(tarPath, ".gz") {
		if fileReader, err = gzip.NewReader(f); err != nil {
			return "", err
		}
	}

	tarReader := tar.NewReader(fileReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if strings.Contains(strings.ToLower(header.Name), "readme") {
			data, err := io.ReadAll(tarReader)
			if err != nil {
				return "", err
			}
			return string(data[:]), nil
		}
	}

	return "", errors.New("plugin.json not found in archive")
}

func SaveLogos(tarPath string, metadata *db.Payload, destDir string) {
	f, err := os.Open(tarPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	var fileReader io.Reader = f
	if strings.HasSuffix(tarPath, ".gz") {
		if fileReader, err = gzip.NewReader(f); err != nil {
			log.Println(err)
			return
		}
	}

	tarReader := tar.NewReader(fileReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return
		}

		if strings.HasSuffix(header.Name, metadata.Info.Logos.Small) {
			if data, err := io.ReadAll(tarReader); err == nil {
				path := filepath.Join(destDir, "logo-small.svg")
				if err := os.WriteFile(path, data, 0644); err != nil {
					log.Println("Failed to save small logo:", err)
				}
			}
		}

		if strings.HasSuffix(header.Name, metadata.Info.Logos.Large) {
			if data, err := io.ReadAll(tarReader); err == nil {
				path := filepath.Join(destDir, "logo-large.svg")
				if err := os.WriteFile(path, data, 0644); err != nil {
					log.Println("Failed to save large logo:", err)
				}
			}
		}
	}
}

func GetLogo(c *gin.Context) {
	slug := c.Param("slug")
	variant := c.Param("variant")

	if variant != "small" && variant != "large" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "variant must be 'small' or 'large'"})
		return
	}

	logoPath := fmt.Sprintf("./static/plugins/%s/logo-%s.svg", slug, variant)

	if _, err := os.Stat(logoPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "logo not found"})
		return
	}

	c.File(logoPath)
}
