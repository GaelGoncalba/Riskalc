package models

import (
	"errors"
)

type Operacion struct {
	ID              string
	Sueldo          float64
	Deudas          float64
	OperacionPagada bool
}

// Validaciones al crear una nueva operación
func NewOperacion(id string, sueldo float64, deudas float64, operacionPagada bool) (Operacion, error) {
	// ID no vacío
	if id == "" {
		return Operacion{}, errors.New("el ID no puede estar vacío")
	}

	// sueldo no sea negativo
	if sueldo < 0 {
		return Operacion{}, errors.New("el sueldo no puede ser negativo")
	}

	// deudas no negativas
	if deudas < 0 {
		return Operacion{}, errors.New("las deudas no pueden ser negativas")
	}

	// Crear y devolver una nueva operación
	return Operacion{
		ID:              id,
		Sueldo:          sueldo,
		Deudas:          deudas,
		OperacionPagada: operacionPagada,
	}, nil
}
