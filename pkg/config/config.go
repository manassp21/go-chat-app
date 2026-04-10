package config

import(
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type Config struct{
	Port                 string
	Host                 string
	DBPath               string
	JWTSecret            string
	JWTExpirationHours   int
	WSReadBufferSize     int
	WSWriteBufferSize    int
}

var AppConfig *Config

func LoadConfig() *Config{
	err:=godotenv.Load()
	if err!=nil{
		log.Println("no .env file found")
	}

	AppConfig = &Config{
		Port:               getEnv("PORT", "8080"),
		Host:               getEnv("HOST", "localhost"),
		DBPath:             getEnv("DB_PATH", "./chat.db"),
		JWTSecret:          getEnv("JWT_SECRET", "your-super-secret-key"),
		JWTExpirationHours: getEnvInt("JWT_EXPIRATION_HOURS", 24),
		WSReadBufferSize:   getEnvInt("WS_READ_BUFFER_SIZE", 1024),
		WSWriteBufferSize:  getEnvInt("WS_WRITE_BUFFER_SIZE", 1024),
	}

	return AppConfig
}

func getEnv(key, defaultVal string) string{
	if val, exists:=os.LookupEnv(key); exists{
		return val
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int{
	valStr:=getEnv(key, "")
	if val, error:=strconv.Atoi(valStr); error==nil{
		return val 
	}
	return defaultVal
}