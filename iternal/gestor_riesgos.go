package gestor

import (
	"errors"
	"fmt"
	"riskalc/internal/models"
	"riskalc/internal/comparador"
)

type GestorRiesgos struct {
	NuevaOperacion        models.Operacion
	Clientes	      map[string]models.Cliente
}

func NuevoGestor(nuevaOperacion models.Operacion, clientes map[string]models.Cliente) *GestorRiesgos {
	return &GestorRiesgos{
		NuevaOperacion:		nuevaOperacion,
		Clientes:  		clientes,
	}, nil
}
