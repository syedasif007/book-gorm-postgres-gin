package setup

import (
	"book-gorm-postgres-gin/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

func SetupModels() *gorm.DB {
	//db, err := gorm.Open("sqlite3", "test.db")
	// Enable VIPER to read Environment Variables
	// viper.AutomaticEnv()

	// To get the value from the config file using key
	// viper package read .env
	// viper_user := viper.Get("POSTGRES_USER")
	// viper_password := viper.Get("POSTGRES_PASSWORD")
	// viper_db := viper.Get("POSTGRES_DB")
	// viper_host := viper.Get("POSTGRES_HOST")
	// viper_port := viper.Get("POSTGRES_PORT")
	// https://gobyexample.com/string-formatting

	// prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	// fmt.Println("conname is\t\t", prosgret_conname)
	// db, err := gorm.Open("postgres", prosgret_conname)

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=123456")
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.Book{})
	// Initialise value
	m := models.Book{Author: "author1", Title: "title1"}
	db.Create(&m)
	return db
}
