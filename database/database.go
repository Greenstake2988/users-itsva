package database

import (
	"users-itsva/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

// Conexion Base de Datos
func (d *DB) ConnectDB() {
	var err error
	d.DB, err = gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("Error al conectar ala base de datos")
	}

	// Migrate the schema
	d.DB.AutoMigrate(&models.User{})

}
