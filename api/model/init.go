package model

import (
	"time"

	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"

	"peckergo/api/config"

	"github.com/jinzhu/gorm"
)

var (
	// DB global gorm instance
	DB         *gorm.DB
	tableCache *cache.Cache
)

// OpenDB open grom
func OpenDB() {
	var err error

	dsn := config.GetString("db.dsn")

	tableCache = cache.New(5*time.Minute, 10*time.Minute)

	log.Info("OpeningDB: ", dsn)
	DB, err = gorm.Open("mysql", dsn)

	DB.LogMode(config.GetBool("db.showlog"))

	if err != nil {
		log.Panic(err.Error())
	}
	Migrate()
	createAdminUser()
}

// HasTable 判断表存在
func HasTable(value interface{}) bool {
	var (
		scope     = DB.NewScope(value)
		tableName string
	)

	if name, ok := value.(string); ok {
		tableName = name
	} else {
		tableName = scope.TableName()
	}
	if _, ok := tableCache.Get(tableName); ok {
		return true
	}
	if DB.HasTable(value) {
		tableCache.Set(tableName, true, cache.NoExpiration)
		return true
	}
	return false
}

// Migrate db migration
func Migrate() {
	log.Info("Migrate ... ")
	DB.AutoMigrate(
		new(User),
		new(LogManagement),
		new(Dbtables),
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
