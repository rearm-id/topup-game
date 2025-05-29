package config

type AppConfig struct {
	Alpine *Alpine
}

func Load() *AppConfig {
	return &AppConfig{}
}
