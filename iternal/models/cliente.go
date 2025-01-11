- cliente.go:
package models

import (
	"errors"
)

type Cliente struct {
	DNI		   string
	OperacionHistorial map[string]Operacion
}

// Validaciones al crear un nuevo cliente
func NewCliente(dni string) (Cliente, error) {
	// ID no vacio
	if dni == "" {
		return Cliente{}, errors.New("el DNI no puede estar vacio")
	}

	// Crear y devolver un nuevo cliente
	return Cliente{
		DNI:                dni,
		OperacionHistorial: make(map[string]Operacion), // Inicializar el historial como map vacio
	}, nil
}
