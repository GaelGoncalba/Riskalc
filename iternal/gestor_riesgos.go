package gestor

import (
	"errors"
	"fmt"
	"riskalc/internal/models"
	"riskalc/internal/comparador"
)

type GestorRiesgos struct {
	Clientes            []models.Cliente
	Operaciones	    []models.Operacion
}

func NuevoGestor(cliente models.Cliente) *GestorRiesgos {
	return &GestorRiesgos{
		Clientes:            []models.Cliente{},
		Operaciones:         []models.Operacion{},
	}
}
