package models

import (
	"errors"
)

type Operacion struct {
	ID             string
	Sueldo         float64
	Deudas         float64
	OperacionPagada bool
}

// Validaciones al crear una nueva operacion
func NewOperacion(id string, sueldo, deudas float64, operacionPagada bool) (Operacion, error) {
	// ID no vacio
	if id == "" {
		return Operacion{}, errors.New("el ID no puede estar vacio")
	}

	// Sueldo no negativo
	if sueldo < 0 {
		return Operacion{}, errors.New("el sueldo no puede ser negativo")
	}

	// Deudas no negativas
	if deudas < 0 {
		return Operacion{}, errors.New("las deudas no pueden ser negativas")
	}

	// Crear y devolver una nueva operacion
	return Operacion{
		ID:             id,
		Sueldo:         sueldo,
		Deudas:         deudas,
		OperacionPagada: operacionPagada,
	}, nil
}
