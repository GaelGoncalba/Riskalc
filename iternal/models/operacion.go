package models

import (
	"errors"
)

type Operacion struct {
	id              string
	sueldo          Sueldo
	deudas          Deuda
	operacionPagada bool
}

// Validaciones al crear una nueva operación
func NewOperacion(id1 string, sueldo1 Sueldo, deudas1 Deuda, operacionPagada1 bool) (Operacion, error) {
	// ID no vacío
	if id1 == "" {
		return Operacion{}, errors.New("el ID no puede estar vacío")
	}

	// Crear y devolver una nueva operación
	return Operacion{
		id:              id1,
		sueldo:          sueldo1,
		deudas:          deudas1,
		operacionPagada: operacionPagada1,
	}, nil
}
