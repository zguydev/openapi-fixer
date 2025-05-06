package config

type FixerConfig struct {
	Tool *ToolConfig `mapstructure:"x-openapi-fixer"`
}

type ToolConfig struct {
	Logger *LoggerConfig `mapstructure:"logger"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}
