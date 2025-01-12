package models

import (
	"errors"
)

type Sueldo struct {
	Valor float64
}

// Validaciones al crear una nueva operación
func NewSueldo(valor float64) (Sueldo, error) {
	// valor no negativo
	if valor < 0 {
		return Sueldo{}, errors.New("el sueldo no puede ser negativo")
	}

	// Crear y devolver el sueldo
	return Operacion{
		Valor: valor,
	}, nil
}
