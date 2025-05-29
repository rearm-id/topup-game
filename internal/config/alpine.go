package config

type Alpine struct {
	Version     string
	AjaxVersion string
}

func (a *AppConfig) GetAlpine() *Alpine {
	return &Alpine{
		Version:     "3.14.1",
		AjaxVersion: "0.12.2",
	}
}
