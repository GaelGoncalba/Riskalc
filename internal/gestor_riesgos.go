package gestor

import (
	"errors"
	"fmt"
	"riskalc/internal/models"
	"riskalc/internal/comparador"
)

type GestorRiesgos struct {
	Clientes            []models.Cliente
}

func NuevoGestor(cliente models.Cliente) *GestorRiesgos {
	return &GestorRiesgos{
		Clientes:            []models.Cliente{},
	}
}
