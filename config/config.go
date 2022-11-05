package config

import (
	"Praktikum/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_HOST     string = "ec2-44-209-57-4.compute-1.amazonaws.com"
	DB_NAME     string = "da9tlh4hfsoqb7"
	DB_USER     string = "dhybjakgyqqmxm"
	DB_PORT     string = "5432"
	DB_PASSWORD string = "c6c3682f069112556a1aca63d6b28b4476eeea6f81ad591a809892559b7589c4"
)

func InitDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	initMigration()
}

func initMigration() {
	DB.AutoMigrate(&model.User{}, &model.Book{})
}
