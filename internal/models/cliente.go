package models

import (
	"errors"
)

type Cliente struct {
	DNI		 string
	Sueldo           float64
	Deudas           float64
	OperacionHistorial []Operacion
}

// Validaciones al crear un nuevo cliente
func NewCliente(dni string, sueldo, deudas float64) (Cliente, error) {
	// ID no vacio
	if dni == "" {
		return Cliente{}, errors.New("el DNI no puede estar vacio")
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
		DNI:              dni,
		Sueldo:           sueldo,
		Deudas:           deudas,
		OperacionHistorial: []Operacion{}, // Inicializar el historial vacio
	}, nil
}
