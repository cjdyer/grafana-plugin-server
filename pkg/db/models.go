package db

type Type string

const (
	TypeDataSource     Type = "datasource"
	TypePanel          Type = "panel"
	TypeApp            Type = "app"
	TypeRenderer       Type = "renderer"
	TypeSecretsManager Type = "secretsmanager"
)

type Plugin struct {
	ID   string `db:"id" json:"id"`
	Type Type   `db:"type" json:"type"`
	Name string `db:"name" json:"name"`
	URL  string `db:"url" json:"url"`
}
