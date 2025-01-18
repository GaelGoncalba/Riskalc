package models

import (
	"errors"
)

type Cliente struct {
	ID		   string
	OperacionHistorial map[string]Operacion
}

// Validaciones al crear un nuevo cliente
func NewCliente(id string) (Cliente, error) {
	// ID no vacio
	if id == "" {
		return Cliente{}, errors.New("el ID no puede estar vacio")
	}

	// Crear y devolver un nuevo cliente
	return Cliente{
		ID:                 id,
		OperacionHistorial: make(map[string]Operacion), // Inicializar el historial como map vacio
	}, nil
}
