package repository_test

import (
	"log"

	"github.com/kunlanat/go-example/migration"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openTestDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Migrate the schema
	migration.AutoMigrate(db)
	return db
}
