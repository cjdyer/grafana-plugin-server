package plugins

import (
	"encoding/json"
	"fmt"

	"github.com/cjdyer/grafana-plugin-server/pkg/db"
)

func AddPlugin(p db.Plugin) error {
	tx := db.DB.MustBegin()

	keywordsJSON, err := json.Marshal(p.Keywords)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO plugins
			(slug, type_id, type_name, type_code, name, url, description, org_name, org_url, keywords, version, updated_at, readme)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.Slug,
		p.TypeId,
		p.TypeName,
		p.TypeCode,
		p.Name,
		p.URL,
		p.Description,
		p.OrgName,
		p.OrgUrl,
		string(keywordsJSON),
		p.Version,
		p.UpdatedAt,
		p.Readme,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func ListPlugins() ([]db.Plugin, error) {
	var plugins []db.Plugin = make([]db.Plugin, 0)
	err := db.DB.Select(&plugins, `
		SELECT id, slug, type_id, type_name, type_code, name, url, description, org_name, org_url, keywords, version, updated_at
		FROM plugins
	`)
	if err != nil {
		return nil, err
	}

	for i := range plugins {
		var kws []string
		if err := json.Unmarshal([]byte(plugins[i].KeywordsJSON), &kws); err == nil {
			plugins[i].Keywords = kws
		} else {
			plugins[i].Keywords = []string{}
		}

		plugins[i].Links = BuildPluginLinks(plugins[i].Slug, plugins[i].Version)
	}

	return plugins, nil
}

func BuildPluginLinks(slug string, version string) []db.Link {
	base := fmt.Sprintf("/api/plugins/%s", slug)
	links := []db.Link{
		{Rel: "self", Href: base},
		{Rel: "versions", Href: base + "/versions"},
		{Rel: "latest", Href: fmt.Sprintf("%s/versions/%s", base, version)},
		{Rel: "download", Href: fmt.Sprintf("%s/versions/%s/download", base, version)},
	}

	return links
}

func UpdatePlugin(p db.Plugin) error {
	tx := db.DB.MustBegin()

	keywordsJSON, err := json.Marshal(p.Keywords)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE plugins
		SET 
			type_id = ?,
			type_name = ?,
			type_code = ?,
			name = ?,
			url = ?,
			description = ?,
			org_name = ?,
			org_url = ?,
			keywords = ?,
			version = ?,
			updated_at = ?
			readme = ?
		WHERE slug = ?
	`,
		p.TypeId,
		p.TypeName,
		p.TypeCode,
		p.Name,
		p.URL,
		p.Description,
		p.OrgName,
		p.OrgUrl,
		string(keywordsJSON),
		p.Version,
		p.UpdatedAt,
		p.Slug,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func GetPluginBySlug(slug string) (*db.Plugin, error) {
	var plugin db.Plugin

	err := db.DB.Get(&plugin, `
		SELECT id, slug, type_id, type_name, type_code, name, url, description, 
		       org_name, org_url, keywords, version, updated_at, readme
		FROM plugins
		WHERE slug = ?
	`, slug)
	if err != nil {
		return nil, err
	}

	var kws []string
	if err := json.Unmarshal([]byte(plugin.KeywordsJSON), &kws); err == nil {
		plugin.Keywords = kws
	} else {
		plugin.Keywords = []string{}
	}

	plugin.Links = BuildPluginLinks(plugin.Slug, plugin.Version)

	return &plugin, nil
}
