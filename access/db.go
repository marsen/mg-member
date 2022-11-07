package access

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"time"
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

func (o *DB) Seed() {
	//todo connection database
	dbURI := "host=localhost user=admin password=1234 dbname=member sslmode=disable" //todo connection string from config
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	m1 := o.NewMember("Mark", 18)
	m2 := o.NewMember("Iris", 29)
	m3 := o.NewMember("Jon", 34)
	db.Table("members").Create(m1)
	db.Table("members").Create(m2)
	db.Table("members").Create(m3)
}

func (o *DB) NewMember(name string, age int) *Member {
	return &Member{
		Model: Model{
			ID:        rand.Int63(),
			CreatedBy: "Seeder",
			CreatedAt: time.Now().UTC(),
			UpdatedBy: "Seeder",
			UpdatedAt: time.Now().UTC(),
			IsValid:   true,
		},
		Age:  age,
		Name: name,
	}
}

type Model struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
	IsValid   bool
}

type Member struct {
	Model
	Name string
	Age  int
}
