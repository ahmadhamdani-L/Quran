package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang-api1/models"

)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME     = "coba1"
)

func ConnectDb() *gorm.DB {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open("postgres", dbinfo)

	checkErr(err)

	db.AutoMigrate(&models.Juz{})
	db.AutoMigrate(&models.Surah{})


	db.AutoMigrate(&models.Juz{}, &models.Surah{})
	db.Model(&models.Juz{}).AddForeignKey("id", "juzs(id)", "CASCADE", "CASCADE")

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
