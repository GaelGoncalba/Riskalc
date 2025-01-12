package gestor

import (
	"errors"
	"fmt"
	"riskalc/internal/models"
	"riskalc/internal/comparador"
)

type GestorRiesgos struct {
	Cliente               models.Cliente
	NuevaOperacion        models.Operacion
	OperacionesHistoricas map[string]map[string]models.Operacion
}

func NuevoGestor(cliente models.Cliente, nuevaOperacion models.Operacion, operacionesHistoricas map[string]map[string]models.Operacion) *GestorRiesgos {
	return &GestorRiesgos{
		Cliente:            	cliente,
		NuevaOperacion:		nuevaOperacion,
		OperacionesHistoricas:  operacionesHistoricas,
	}, nil
}
