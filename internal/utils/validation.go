package utils

import (
	"errors"
	"regexp"
	"strings"
)

func NameValidator(name string) error {
	re := regexp.MustCompile(`^[A-Za-z0-9 ,.]+$`)
	if !re.MatchString(name) {
		return errors.New("name only allows letters, numbers, spaces, commas, and dots")
	}
	return nil
}

func LanguageValidator(lan string) error {
	re := regexp.MustCompile(`^[A-Za-z]+$`)
	if !re.MatchString(lan) {
		return errors.New("the language must only contain letters (no spaces or symbols)")
	}
	return nil
}

func CategoryValidator(category string) error {
	re := regexp.MustCompile(`^[A-Za-z ]+$`)
	if !re.MatchString(category) {
		return errors.New("the category only allows letters and spaces")
	}
	return nil
}

func DescriptionValidator(description string) error {
	re := regexp.MustCompile(`^[A-Za-z0-9 ,.]+$`)
	if !re.MatchString(description) {
		return errors.New("the description only allows letters, numbers, spaces, commas, and periods")
	}
	return nil
}

func CodeValidator(code string) error {
	if strings.TrimSpace(code) == "" {
		return errors.New("the code cannot be empty")
	}
	return nil
}

func SanitizeInput(input string) string {
	return strings.TrimSpace(input)
}
