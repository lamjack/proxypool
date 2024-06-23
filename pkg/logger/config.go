package logger

const (
	DevelopmentLevel = "development"
	ProductionLevel  = "production"

	defaultLevel = DevelopmentLevel
)

type Config struct {
	Level string `json:"level"`
}

func (l *Config) GetLevel() string {
	if l.Level == "" {
		return defaultLevel
	}
	return l.Level
}
