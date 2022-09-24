package models

type Settings struct {
	AppParams      Params           `json:"app"`
	PostgresParams PostgresSettings `json:"postgres_params"`
}

type Params struct {
	ServerURL     string `json:"server_url"`
	ServerName    string `json:"server_name"`
	AppVersion    string `json:"app_version"`
	PortRun       string `json:"port_run"`
	LogInfo       string `json:"log_info"`
	LogError      string `json:"log_error"`
	LogDebug      string `json:"log_debug"`
	LogWarning    string `json:"log_warning"`
	LogMaxSize    int    `json:"log_max_size"`
	LogMaxBackups int    `json:"log_max_backups"`
	LogMaxAge     int    `json:"log_max_age"`
	LogCompress   bool   `json:"log_compress"`
}

type PostgresSettings struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}
