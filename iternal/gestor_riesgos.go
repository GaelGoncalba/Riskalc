package gestor

import (
	"errors"
	"fmt"
	"riskalc/internal/models"
	"riskalc/internal/comparador"
)

type GestorRiesgos struct {
	Cliente             models.Cliente
	NuevaOperacion      models.Operacion
	OperacionesHistoricas []models.Operacion
}

func NuevoGestor(cliente models.Cliente, nuevaOperacion models.Operacion, operacionesHistoricas []models.Operacion) *GestorRiesgos {
	return &GestorRiesgos{
		Cliente:            	cliente,
		NuevaOperacion:		nuevaOperacion,
		OperacionesHistoricas:  operacionesHistoricas,
	}, nil
}
