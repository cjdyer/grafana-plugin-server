package plugins

import (
	"log"

	"github.com/cjdyer/grafana-plugin-server/pkg/db"
)

type PluginWithVersions struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"`
	Versions []db.Version `json:"versions"`
}

func ListPluginsWithVersions() ([]PluginWithVersions, error) {
	rows, err := db.DB.Query(`SELECT id, type FROM plugins`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PluginWithVersions
	for rows.Next() {
		var p db.Plugin
		if err := rows.Scan(&p.ID, &p.Type); err != nil {
			return nil, err
		}

		// fetch versions for each plugin
		vrows, err := db.DB.Query(`SELECT id, plugin_id, version, url FROM versions WHERE plugin_id = ?`, p.ID)
		if err != nil {
			return nil, err
		}

		var versions []db.Version
		for vrows.Next() {
			var v db.Version
			if err := vrows.Scan(&v.ID, &v.PluginID, &v.Version, &v.URL); err != nil {
				return nil, err
			}
			versions = append(versions, v)
		}
		vrows.Close()

		log.Default().Println(p)

		results = append(results, PluginWithVersions{
			ID:       p.ID,
			Type:     p.Type,
			Versions: versions,
		})
	}

	return results, nil
}

func AddPluginWithVersion(p db.Plugin, v db.Version) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	// Insert plugin if not exists
	_, err = tx.Exec(`INSERT OR IGNORE INTO plugins (id, type) VALUES (?, ?)`, p.ID, p.Type)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert version
	_, err = tx.Exec(
		`INSERT INTO versions (plugin_id, version, url) VALUES (?, ?, ?)`,
		v.PluginID, v.Version, v.URL,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
