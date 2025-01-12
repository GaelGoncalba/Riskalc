package models

import (
	"errors"
)

type Operacion struct {
	ID              string
	Sueldo          Sueldo
	Deudas          Deuda
	OperacionPagada bool
}

// Validaciones al crear una nueva operación
func NewOperacion(id string, sueldo Sueldo, deudas Deuda, operacionPagada bool) (Operacion, error) {
	// ID no vacío
	if id == "" {
		return Operacion{}, errors.New("el ID no puede estar vacío")
	}

	// Crear y devolver una nueva operación
	return Operacion{
		ID:              id,
		Sueldo:          sueldo,
		Deudas:          deudas,
		OperacionPagada: operacionPagada,
	}, nil
}
