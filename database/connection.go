package database

import (
	"log"
	"seleksi-compfest-backend/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return err
	}
	log.Println("Connected to database")
	return nil
}

func MigrateProduct(table *entity.Product) {
	Connector.AutoMigrate(&table)
	Connector.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&table)
	log.Println("Migrated table product")
}
