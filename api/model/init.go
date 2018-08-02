package model

import (
	log "github.com/sirupsen/logrus"

	"peckergo/api/config"

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
		//ph-AutoMigrate don't remove this line
	)
}

// CloseDB close gorm
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func createAdminUser() {
	u := &User{
		DisplayName: "admin",
		Username:    config.GetString("admin.userName"),
		Password:    config.GetString("admin.password"),
		Role:        RoleAdmin,
	}
	NewUser(u)
}
