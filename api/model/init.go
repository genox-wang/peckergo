package model

import (
	"console-template/api/utils/log"

	"console-template/api/config"

	"github.com/jinzhu/gorm"
)

var (
	// DB global gorm instance
	DB *gorm.DB
)

// OpenDB open grom
func OpenDB() {
	var err error

	dsn := config.GetString("db.dsn")

	log.Info("OpeningDB: ", dsn)
	DB, err = gorm.Open("mysql", dsn)

	DB.LogMode(config.GetBool("db.showlog"))

	if err != nil {
		log.Panic(err.Error())
	}
	Migrate()
	createAdminUser()
}

// Migrate db migration
func Migrate() {
	log.Info("Migrate ... ")
	DB.AutoMigrate(
		new(User),
	)
}

// CloseDB close gorm
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func createDatabase() error {
	db, err := gorm.Open("mysql", config.GetString("db.mysql"))
	if err == nil {
		return db.Exec("CREATE DATABASE `console-template` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;").Error
	}
	return err
}

func createAdminUser() {
	u := &User{
		DisplayName: "admin",
		Username:    config.GetString("admin.userName"),
		Password:    config.GetString("admin.password"),
	}
	NewUser(u)
}
