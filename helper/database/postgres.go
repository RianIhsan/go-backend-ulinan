package database

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ulinan/config"
	"ulinan/entities"
)

func BootDatabase(cnf config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cnf.Database.DbHost, cnf.Database.DbUser, cnf.Database.DbPass, cnf.Database.DbName, cnf.Database.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database", err.Error())
		return nil
	}
	log.Info("Database connected")
	return db
}

func MigrateTable(db *gorm.DB) {
	err := db.AutoMigrate(
		entities.UserEntity{},
		entities.GenderEntity{},
	)
	if err != nil {
		log.Fatal("Failed to migrate table", err.Error())
	} else {
		log.Info("Table successfully migrated")
	}
}
