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

// === DATABASE ===

// DbHost ... Database connection ip
var DbHost = os.Getenv("DB_HOST")

// DbPort ... Database connection port
var DbPort = os.Getenv("DB_PORT")

// DbSSL ... Database connection ssl mode
var DbSSL = os.Getenv("DB_SSL")

// DbDatabase ... Database connection main database
var DbDatabase = os.Getenv("DB_DATABASE")

// DbUsername ... Database connection username
var DbUsername = os.Getenv("DB_USERNAME")

// DbPassword ... Database connection password
var DbPassword = os.Getenv("DB_PASSWORD")

// === CACHE ===

// RedisHost ... Redis connection host
var RedisHost = os.Getenv("REDIS_HOST")

// RedisPort ... Redis connection port
var RedisPort = os.Getenv("REDIS_PORT")

// RedisPassword ... Redis connection password
var RedisPassword = os.Getenv("REDIS_PASSWORD")
