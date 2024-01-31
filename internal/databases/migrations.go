package databases

import (
	"log"

	"github.com/kkrajkumar1198/blog-grpc/internal/blog/models"
)

type Migrations struct {
	DB IConnection
}

func (m Migrations) MigrateData() {
	connection, conErr := m.DB.GetConnection()

	if conErr != nil {
		log.Println("Connection Error Migrations: ", conErr.Error())
	}

	connection.AutoMigrate(&models.Post{})
}

func InitMigrations() {
	migrations := Migrations{
		DB: SQLiteDB{},
	}

	migrations.MigrateData()
}
