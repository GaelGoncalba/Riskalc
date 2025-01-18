package models

import (
	"errors"
)

type Cliente struct {
	id		   string
	operacionHistorial map[string]Operacion
}

// Validaciones al crear un nuevo cliente
func NewCliente(id1 string) (Cliente, error) {
	// ID no vacio
	if id1 == "" {
		return Cliente{}, errors.New("el ID no puede estar vacio")
	}

	// Crear y devolver un nuevo cliente
	return Cliente{
		id:                 id1,
		operacionHistorial: make(map[string]Operacion), // Inicializar el historial como map vacio
	}, nil
}
