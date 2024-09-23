package migration

import (
	"fmt"
	"log"

	"github.com/kunlanat/go-example/repository/entities"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if db == nil {
		log.Fatal("DB connection can't nill")
	}

	migrates := []interface{}{
		&entities.Books{},
	}

	if err := db.AutoMigrate(migrates...); err != nil {
		log.Fatal(fmt.Sprintf("Can't automigrate schema %v", err))
	}
}
