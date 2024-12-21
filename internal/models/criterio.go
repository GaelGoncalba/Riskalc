package models

import (
	"errors"
)

type Criterio struct {
	SueldoMinimo    float64
	DeudaMaxima     float64
	OperacionPagada bool
}

// Validaciones al crear un nuevo criterio
func NewCriterio(sueldoMinimo, deudaMaxima float64, operacionPagada bool) (Criterio, error) {
	// Sueldo minimo no negativo
	if sueldoMinimo < 0 {
		return Criterio{}, errors.New("el sueldo minimo no puede ser negativo")
	}

	// Deuda maxima no negativa
	if deudaMaxima < 0 {
		return Criterio{}, errors.New("la deuda maxima no puede ser negativa")
	}

	// Crear y devolver un nuevo criterio
	return Criterio{
		SueldoMinimo:    sueldoMinimo,
		DeudaMaxima:     deudaMaxima,
		OperacionPagada: operacionPagada,
	}, nil
}
