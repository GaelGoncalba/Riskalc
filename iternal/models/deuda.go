package models

import (
	      "errors"
)

type Deuda struct {
	valor float64
}

// Validaciones al crear una nueva deuda
func NuevaDeuda(valor1 float64) (Deuda, error) {
  // deuda no negativa
	if valor1 < 0 {
		return Deuda{}, errors.New("la deuda no puede ser negativa")
	}
  
	return Deuda{
    valor: valor1
  }, nil
}
