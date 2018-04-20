package config

type Configuration struct {
	AppConfig *AppConfig `toml:"App"`
}

type AppConfig struct {
	Address string `toml:"Address"`
	RunMode string `toml:"RunMode"`
}

func NewDefaultConfiguration() *Configuration {
	return &Configuration{
		AppConfig: &AppConfig{
			Address: "0.0.0.0:8080",
			RunMode: "stage",
		},
	}
}
