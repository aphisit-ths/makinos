package models

type AppSettings struct {
	MakinoConfiguration struct {
		AmeliaApiUrl   string `json:"AmeliaApiUrl"`
		VegapunkApiUrl string `json:"VegapunkApiUrl"`
	} `json:"MakinoConfiguration"`
	ConnectionStrings struct {
		CustomerDocumentDb   string `json:"CustomerDocumentDb"`
		PostgresDb           string `json:"PostgresDb"`
		CustomerDatabaseName string `json:"CustomerDatabaseName"`
	} `json:"ConnectionStrings"`
}
