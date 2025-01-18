package models

import (
	"errors"
)

type Sueldo struct {
	valor float64
}

// Validaciones al crear una nueva operación
func NewSueldo(valor1 float64) (Sueldo, error) {
	// valor no negativo
	if valor1 < 0 {
		return Sueldo{}, errors.New("el sueldo no puede ser negativo")
	}

	// Crear y devolver el sueldo
	return Operacion{
		valor: valor1,
	}, nil
}
