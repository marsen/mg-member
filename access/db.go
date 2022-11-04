package access

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
}

func (o *DB) Migrate() {
	//todo connection database
	dbURI := "host=localhost user=admin password=1234 dbname=member sslmode=disable" //todo connection string from config
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Member{})
	if err != nil {
		panic(err)
	}
}

type Member struct {
	gorm.Model
	Name string
	Age  int
}
