package repository

import (
	"fmt"
	"strconv"

	"github.com/patosinsuerte/ducky-clip/internal/database"
	"github.com/patosinsuerte/ducky-clip/models"
)

func ListAll() ([]models.Snippet, error) {
	var snippets []models.Snippet
	result := database.DB.Find(&snippets)
	if result.Error != nil {
		return nil, result.Error
	}
	return snippets, nil
}

func AddOne(snippet *models.Snippet) error {
	result := database.DB.Create(snippet)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteAll() error {
	result := database.DB.Where("1 = 1").Delete(&models.Snippet{})

	if result.Error != nil {
		return result.Error
	}
	// fmt.Printf("Se han vaciado %d snippets del estanque. 🦆🌊\n", result.RowsAffected)
	return nil
}

func ListById(id string) (*models.Snippet, error) {

	u64, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("el ID proporcionado no es válido: %v", err)
	}

	var snippet models.Snippet

	result := database.DB.First(&snippet, u64)

	if result.Error != nil {
		return nil, result.Error
	}

	return &snippet, nil
}

func DeleteById(id string) error {
	// se convierte el id que viene de string primero
	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fmt.Errorf("el ID '%s' no es un número válido", id)
	}
	result := database.DB.Delete(&models.Snippet{}, u64)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún snippet con el ID %d", u64)
	}

	return nil
}
