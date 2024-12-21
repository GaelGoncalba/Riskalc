package models

import (
	"errors"
)

type Cliente struct {
	ID               string
	Sueldo           float64
	Deudas           float64
	OperacionHistorial []Operacion
}

// Validaciones al crear un nuevo cliente
func NewCliente(id string, sueldo, deudas float64) (Cliente, error) {
	// ID no vacio
	if id == "" {
		return Cliente{}, errors.New("el ID no puede estar vacio")
	}

	// Sueldo no negativo
	if sueldo < 0 {
		return Cliente{}, errors.New("el sueldo no puede ser negativo")
	}

	// Deudas no negativas
	if deudas < 0 {
		return Cliente{}, errors.New("las deudas no pueden ser negativas")
	}

	// Crear y devolver un nuevo cliente
	return Cliente{
		ID:               id,
		Sueldo:           sueldo,
		Deudas:           deudas,
		OperacionHistorial: []Operacion{}, // Inicializar el historial vacio
	}, nil
}
