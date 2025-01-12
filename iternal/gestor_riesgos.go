package gestor

import (
	"errors"
	"fmt"
	"riskalc/internal/models"
	"riskalc/internal/comparador"
)

type GestorRiesgos struct {
	NuevaOperacion        models.Operacion
	OperacionesHistoricas map[string]models.Cliente
}

func NuevoGestor(nuevaOperacion models.Operacion, operacionesHistoricas map[string]models.Cliente) *GestorRiesgos {
	return &GestorRiesgos{
		NuevaOperacion:		nuevaOperacion,
		OperacionesHistoricas:  operacionesHistoricas,
	}, nil
}
