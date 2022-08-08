package postgres

import (
	"fmt"
	"log"
	"strconv"

	"gitlab.com/cinco/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/schema"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := configs.Config().Dbconfig.Port
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("Error when parsing port environment!")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configs.Config().Dbconfig.Host, port, configs.Config().Dbconfig.Username, configs.Config().Dbconfig.Password, configs.Config().Dbconfig.Dbname)
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	//err = DB.AutoMigrate(&model.Account{}, &model.User{}, &model.Cashflow{})
	//if err != nil {
	//	panic("[Gorm] Database failed to migrate!")
	//}
	//
	//fmt.Println("Database Migrated")
}
