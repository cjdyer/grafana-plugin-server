package db

type Plugin struct {
	ID   string `db:"id" json:"id"`
	Type string `db:"type" json:"type"`
}

type Version struct {
	ID       int    `db:"id" json:"id"`
	PluginID string `db:"plugin_id" json:"plugin_id"`
	Version  string `db:"version" json:"version"`
	URL      string `db:"url" json:"url"`
}
