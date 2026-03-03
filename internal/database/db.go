package database

import (
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// func InitDB() error {
// 	path, err := GetDBPath()
// 	if err != nil {
// 		return err
// 	}

// 	//Conexión con GORM
// 	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}

// 	//se crea la tabla con SQL
// 	/*
// 			Esto seria con mas codigo asi:
// 			err := createDBTables(db) // Creamos la variable afuera
// 			if err != nil {
// 		    	return err
// 			}
// 			// Aquí la variable 'err' sigue ocupando espacio en la memoria mientras que con esa sintaxis NO
// 	*/
// 	if err := createDBTables(db); err != nil {
// 		return err
// 	}

// 	DB = db
// 	return nil
// }

// CREA LA CARPETA .ducky-clip como oculta en la carpeta del usuario
func GetDBPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	var baseDir string
	if err != nil {
		baseDir, _ = os.Getwd()
	} else {
		baseDir = filepath.Join(homeDir, ".ducky-clip")
	}
	os.MkdirAll(baseDir, 0755)
	return filepath.Join(baseDir, "snippets.db"), nil
}

func createDBTables(db *gorm.DB) error {

	query := `
	CREATE TABLE IF NOT EXISTS Snippets(
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name TEXT NOT NULL CHECK (name NOT GLOB '*[^A-Za-z0-9 ,.]*'),
    	language TEXT NOT NULL CHECK (language NOT GLOB '*[^A-Za-z]*'),
    	code TEXT NOT NULL,
    	category TEXT CHECK (category NOT GLOB '*[^A-Za-z ]*'),
    	description TEXT CHECK (description NOT GLOB '*[^A-Za-z0-9 ,.]*'),
    	CONSTRAINT code_not_empty CHECK (length(code) > 0),
    	CONSTRAINT name_not_empty CHECK (length(name) > 0),
    	CONSTRAINT language_not_empty CHECK (length(language) > 0)
	);`
	return db.Exec(query).Error
}

func InitDB() error {
	path, err := GetDBPath()
	if err != nil {
		return err
	}

	// Configuración para silenciar los logs técnicos
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // <--- ESTO apaga los mensajes de SLOW SQL
	}

	// 1. Conexión con la configuración silenciosa
	db, err := gorm.Open(sqlite.Open(path), config)
	if err != nil {
		return err
	}

	if err := createDBTables(db); err != nil {
		return err
	}

	DB = db
	return nil
}
