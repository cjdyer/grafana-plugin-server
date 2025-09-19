package db

type TypeCode string

const (
	TypeCodeDataSource     TypeCode = "datasource"
	TypeCodePanel          TypeCode = "panel"
	TypeCodeApp            TypeCode = "app"
	TypeCodeRenderer       TypeCode = "renderer"
	TypeCodeSecretsManager TypeCode = "secretsmanager"
)

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Plugin struct {
	ID           uint16   `db:"id" json:"id"`
	Slug         string   `db:"slug" json:"slug"`
	TypeId       uint8    `db:"type_id" json:"typeId"`
	TypeName     string   `db:"type_name" json:"typeName"`
	TypeCode     TypeCode `db:"type_code" json:"typeCode"`
	Name         string   `db:"name" json:"name"`
	URL          string   `db:"url" json:"url"`
	Description  string   `db:"description" json:"description"`
	OrgName      string   `db:"org_name" json:"orgName"`
	OrgUrl       string   `db:"org_url" json:"orgUrl"`
	Keywords     []string `db:"-" json:"keywords"`
	KeywordsJSON string   `db:"keywords" json:"-"`
	Version      string   `db:"version" json:"version"`
	UpdatedAt    string   `db:"updated_at" json:"updatedAt"`
	Links        []Link   `db:"-" json:"links"`
	FilePath     string   `db:"file_path" json:"filePath"`
}

type Payload struct {
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
		Logos    struct {
			Small string `json:"small"`
			Large string `json:"large"`
		} `json:"logos"`
	} `json:"info"`
}
