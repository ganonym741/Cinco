package configs

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type ConfDB struct {
	Host     string
	Port     string
	Dbname   string
	Username string
	Password string
}

type ConfJWT struct {
	Secret  string
	Expired int
}

type Configs struct {
	Dbconfig  ConfDB
	Jwtconfig ConfJWT
}

var configs *Configs
var lock = &sync.Mutex{}

// Config func to get env value
func Config() *Configs {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	if configs == nil {
		lock.Lock()
		JwtExpired, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_SECOND"))

		configs = &Configs{
			Dbconfig: ConfDB{
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Dbname:   os.Getenv("DB_NAME"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
			},
			Jwtconfig: ConfJWT{
				Secret:  os.Getenv("JWT_SECRET"),
				Expired: JwtExpired,
			},
		}
		lock.Unlock()
	}
	return configs

	// // load .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Print("Error loading .env file")
	// }
	// // Return the value of the variable
	// return os.Getenv(key)
}
