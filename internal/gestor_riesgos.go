package gestor

import (
	"errors"
	"fmt"
	"riskalc/internal/models"
	"riskalc/internal/comparador"
)

type GestorRiesgos struct {
	CriteriosEvaluacion models.Criterio
	Clientes            []models.Cliente
}

func NuevoGestor(criterio models.Criterio, cliente models.Cliente) *GestorRiesgos {
	return &GestorRiesgos{
		CriteriosEvaluacion: criterio,
		Clientes:            []models.Cliente{},
	}
}
