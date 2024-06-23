package configs

import "github.com/spf13/viper"

type Configs struct {
	Db        Db        `json:"db"`
	KuaiDaiLi KuaiDaiLi `json:"kuaidaili"`
}

type Db struct {
	Dsn string `json:"dsn"`
}

type KuaiDaiLi struct {
	SecretId  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
	Num       int    `json:"num"`
}

// NewConfigs creates a new Configs instance by reading and unmarshalling the configuration file.
// It returns a pointer to the created Configs instance and any error encountered.
func NewConfigs() (*Configs, error) {
	viper.SetConfigName("configs")
	viper.SetConfigType("json")

	viper.SetDefault("kuaidaili.num", 1)

	// Add paths where the configuration file could be located.
	// The paths are checked in reverse order, with the last added path having the highest priority.
	viper.AddConfigPath("/etc/proxypool")
	viper.AddConfigPath("$HOME/.proxypool")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// Enable automatic environment variable configuration.
	// Only environment variables with the prefix "proxypool" are considered.
	viper.AutomaticEnv()
	viper.SetEnvPrefix("proxypool")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var configs Configs
	err = viper.Unmarshal(&configs)
	if err != nil {
		return nil, err
	}

	return &configs, nil
}
