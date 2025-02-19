package configs

import "github.com/spf13/viper"

type DocsConfig struct {
	SwaggerUrl string
}

func NewDocsConfig(v *viper.Viper) *DocsConfig {
	return &DocsConfig{
		SwaggerUrl: v.GetString("SWAGGER_URL"),
	}
}
