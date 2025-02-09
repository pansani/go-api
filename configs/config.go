package configs

import (
	jwtAuth "github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type config struct {
	DbHost        string `mapstructure:"DB_HOST"`
	DbPort        string `mapstructure:"DB_PORT"`
	DbUser        string `mapstructure:"DB_USER"`
	DbPassword    string `mapstructure:"DB_PASSWORD"`
	DbName        string `mapstructure:"DB_NAME"`
	DbDriver      string `mapstructure:"DB_DRIVER"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JwtSecret     string `mapstructure:"JWT_SECRET"`
	JwtExpiry     string `mapstructure:"JWT_EXPIRY"`
	TokenAuth     *jwtAuth.JWTAuth
}

func LoadConfig(configFile string) (*config, error) {
	var cfg *config

	viper.SetConfigFile("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(configFile)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtAuth.New("HS256", []byte(cfg.JwtSecret), nil)
	return cfg, nil
}
