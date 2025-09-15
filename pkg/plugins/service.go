package plugins

import "github.com/cjdyer/grafana-plugin-server/pkg/db"

func ListPlugins() ([]db.Plugin, error) {
	return []db.Plugin{}, nil
}

func AddPlugin(p db.Plugin) error {
	return nil
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
