package hexagonal_test

import (
	"strings"
	"testing"

	"github.com/ANDERSON1808/archTest/hexagonal"
)

func TestHexagonalValidator_Validate(t *testing.T) {
	// Crear instancia del validador
	validator := hexagonal.NewHexagonalValidator()

	// Definir los paquetes y sus dependencias permitidas
	validator.AddPackage("Delivery", []string{"UseCase"})
	validator.AddPackage("Repository", []string{"UseCase"})
	validator.AddPackage("UseCase", []string{"Delivery", "Repository"})

	// Definir las dependencias actuales
	dependencies := map[string][]string{
		"Delivery":   {"UseCase"},
		"Repository": {"UseCase"},
		"UseCase":    {"Delivery", "Repository"},
	}

	// Validar las dependencias
	err := validator.Validate(dependencies)

	// Verificar si se generó un error
	if err != nil {
		t.Errorf("Se esperaba que las dependencias sean válidas, pero se encontró el siguiente error: %s", err)
	}
}

func TestHexagonalValidator_Validate_InvalidDependencies(t *testing.T) {
	// Crear instancia del validador
	validator := hexagonal.NewHexagonalValidator()

	// Definir los paquetes y sus dependencias permitidas
	validator.AddPackage("Delivery", []string{"UseCase"})
	validator.AddPackage("Repository", []string{"UseCase"})
	validator.AddPackage("UseCase", []string{"Delivery", "Repository"})

	// Definir las dependencias actuales (con una dependencia no permitida)
	dependencies := map[string][]string{
		"Delivery":   {"UseCase"},
		"Repository": {"UseCase"},
		"UseCase":    {"Delivery", "InvalidDependency"},
	}

	// Validar las dependencias
	err := validator.Validate(dependencies)

	// Verificar si se generó un error
	if err == nil {
		t.Error("Se esperaba un error debido a una dependencia no permitida, pero no se encontró ningún error")
	} else {
		// Verificar el mensaje de error específico
		expectedError := "Invalid dependency: UseCase -> InvalidDependency"
		if !strings.EqualFold(err.Error(), expectedError) {
			t.Errorf("Se esperaba el siguiente error: %s, pero se encontró: %s", expectedError, err.Error())
		}
	}
}

func TestHexagonalValidator_Validate_MissingPackage(t *testing.T) {
	// Crear instancia del validador
	validator := hexagonal.NewHexagonalValidator()

	// Definir los paquetes y sus dependencias permitidas
	validator.AddPackage("Delivery", []string{"UseCase"})
	validator.AddPackage("Repository", []string{"UseCase"})
	validator.AddPackage("UseCase", []string{"Delivery", "Repository"})

	// Definir las dependencias actuales (falta un paquete)
	dependencies := map[string][]string{
		"Delivery":   {"UseCase"},
		"Repository": {"UseCase"},
		// Faltando el paquete "UseCase"
	}

	// Validar las dependencias
	err := validator.Validate(dependencies)

	// Verificar si se generó un error
	if err == nil {
		t.Error("Se esperaba un error debido a un paquete faltante, pero no se encontró ningún error")
	} else {
		// Verificar el mensaje de error específico
		expectedError := "missing package: UseCase"
		if err.Error() != expectedError {
			t.Errorf("Se esperaba el siguiente error: %s, pero se encontró: %s", expectedError, err.Error())
		}
	}
}
