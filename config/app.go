package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Apply config settings from the env, and config file(s
func Apply() {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("config file not found, using defaults")
	}

	setDefaults()

	gin.SetMode(viper.GetString("GIN_MODE"))
}

// Default config settings if not found in the env or config file(s)
func setDefaults() {
	viper.SetDefault("APP_KEY", "")
	viper.SetDefault("APP_ENV", "production")
	viper.SetDefault("APP_DEBUG", false)
	viper.SetDefault("APP_DOMAIN", "")
	viper.SetDefault("ASSET_PATH", "/static")
	viper.SetDefault("ASSET_DIR", "public")
	viper.SetDefault("VIEW_DIR", "views")
	viper.SetDefault("VIEW_EXT", ".jade")
	viper.SetDefault("SESSION_KEY", "pulse_session")
	viper.SetDefault("DATABASE_HOST", "localhost:27017")
	viper.SetDefault("DATABASE_USER", "")
	viper.SetDefault("DATABASE_PASSWORD", "")
	viper.SetDefault("DATABASE_DEFAULT", "pulse")
	viper.SetDefault("REDIS_HOST", "localhost:6379")
	viper.SetDefault("REDIS_PASS", "")
	viper.SetDefault("GIN_MODE", "release")
	viper.SetDefault("PORT", 8080)
}
