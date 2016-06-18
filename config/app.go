package config

import (
	"os"

	// Autoload all the values from .env
	_ "github.com/joho/godotenv/autoload"
)

// === APPLICATION ===

// AppPort ... Application port to run server
var AppPort = os.Getenv("APP_PORT")

// AppEnv ... Application environment
var AppEnv = os.Getenv("APP_ENV")

// AppDebug ... Application debug mode
var AppDebug = os.Getenv("APP_DEBUG")

// AppKey ... Application secret key for signing requests and cookies
var AppKey = os.Getenv("APP_KEY")

// AppDomain ... Application domain for cookie and url generation
var AppDomain = os.Getenv("APP_DOMAIN")

// SessionKey ... Application user session key
var SessionKey = os.Getenv("SESSION_KEY")

// AssetPath ... Application assets url path
var AssetPath = os.Getenv("ASSET_PATH")

// AssetDir ... Application assets location
var AssetDir = os.Getenv("ASSET_DIR")

// ViewDir ... View templates location
var ViewDir = os.Getenv("VIEW_DIR")

// ViewExt ... View file extension
var ViewExt = os.Getenv("VIEW_EXT")

// === DATABASE ===

// DbURL ... Fully qualified database connection url
var DbURL = os.Getenv("DATABASE_URL")

// DbType ... Database connection type
var DbType = os.Getenv("DATABASE_TYPE")

// === CACHE ===

// RedisHost ... Redis connection host
var RedisHost = os.Getenv("REDIS_HOST")

// RedisPassword ... Redis connection password
var RedisPassword = os.Getenv("REDIS_PASSWORD")
