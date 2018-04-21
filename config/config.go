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
		},
	}
}
