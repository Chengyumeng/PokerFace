package config

type Configuration struct {
	AppConfig *AppConfig `toml:"App"`
}

type AppConfig struct {
	Address  string    `toml:"Address"`
	RunMode  string    `toml:"RunMode"`
	Database *Database `toml:"Database"`
}

type Database struct {
	DriverName string `toml:"driver"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	Charset    string `toml:"charset"`
	Tns        string `toml:"tns"`
	ShowSQL    bool   `toml:"showSQL"`
}

func NewDefaultConfiguration() *Configuration {
	return &Configuration{
		AppConfig: &AppConfig{
			Address: "0.0.0.0:8080",
			RunMode: "stage",
			Database: &Database{
				DriverName: "mysql",
				User:       "root",
				Password:   "0a1b2c3d4e",
				Charset:    "utf8",
				Tns:        "tcp(127.0.0.1:4540)/porkerface",
				ShowSQL:    false,
			},
		},
	}
}
