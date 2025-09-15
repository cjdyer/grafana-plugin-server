package plugins

import (
	"github.com/cjdyer/grafana-plugin-server/pkg/db"
)

func ListPlugins() ([]db.Plugin, error) {
	var plugins []db.Plugin = make([]db.Plugin, 0)
	err := db.DB.Select(&plugins, `SELECT id, type, name, url FROM plugins`)
	if err != nil {
		return nil, err
	}
	return plugins, nil
}

func AddPlugin(p db.Plugin) error {
	tx := db.DB.MustBegin()
	tx.Exec(`INSERT INTO plugins (id, type, name, url) VALUES (?, ?, ?, ?)`, p.ID, p.Type, p.Name, p.URL)
	return tx.Commit()
}
