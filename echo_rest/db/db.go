package db

import (
	config "echo_rest/config"
	"echo_rest/migration"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetMYSQLConnectionString()
	log.Print(conString)
	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}
	DB.AutoMigrate(&migration.Pegawai{}, &migration.User{})
	if err := DB.Model(&migration.Pegawai{}).AddUniqueIndex("idx_translations_key_with_locale", "locale", "key").Error; err != nil {
		fmt.Printf("Failed to create unique index for translations key & locale, got: %v\n", err.Error())
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
