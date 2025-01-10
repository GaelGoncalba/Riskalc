- cliente.go:
package models

import (
	"errors"
)

type Cliente struct {
	DNI		 string
	OperacionHistorial []Operacion
}

// Validaciones al crear un nuevo cliente
func NewCliente(dni string) (Cliente, error) {
	// ID no vacio
	if dni == "" {
		return Cliente{}, errors.New("el DNI no puede estar vacio")
	}

	// Crear y devolver un nuevo cliente
	return Cliente{
		DNI:              dni,
		OperacionHistorial: []Operacion{}, // Inicializar el historial vacio
	}, nil
}
