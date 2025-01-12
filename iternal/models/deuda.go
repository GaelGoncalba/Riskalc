package models

import (
	      "errors"
)

type Deuda struct {
	Valor float64
}

// Validaciones al crear una nueva deuda
func NuevaDeuda(valor float64) (Deuda, error) {
  // deuda no negativa
	if valor < 0 {
		return Deuda{}, errors.New("la deuda no puede ser negativa")
	}
  
	return Deuda{
    Valor: valor
  }, nil
}
