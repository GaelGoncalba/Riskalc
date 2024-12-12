package gestor

import (
	"fmt"
	"riskalc/internal/models"
	"riskalc/internal/comparador"
)

type GestorRiesgos struct {
	CriteriosEvaluacion models.Criterio
	Cliente             models.Cliente
}

func NuevoGestor(criterio models.Criterio, cliente models.Cliente) *GestorRiesgos {
	return &GestorRiesgos{
		CriteriosEvaluacion: criterio,
		Cliente:             cliente,
	}
}

func (g *GestorRiesgos) EvaluarRiesgo() string {
	coincidencias := comparador.CompararOperaciones(g.Cliente.OperacionHistorial[0], g.Cliente.OperacionHistorial, g.CriteriosEvaluacion)

	if coincidencias > 3 {
		return fmt.Sprintf("Riesgo bajo: El cliente tiene %d operaciones similares aprobadas.", coincidencias)
	} else if coincidencias > 0 {
		return fmt.Sprintf("Riesgo moderado: El cliente tiene %d operaciones similares.", coincidencias)
	} else {
		return "Riesgo alto: No se encontraron operaciones similares."
	}
}