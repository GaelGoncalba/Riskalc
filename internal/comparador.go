package comparador

import (
	"riskalc/internal/models"
)

func CompararOperaciones(nuevaOperacion models.Operacion, historial []models.Operacion, criterio models.Criterio) int {
	coincidencias := 0

	for _, operacion := range historial {
		if operacion.Sueldo >= criterio.SueldoMinimo &&
			operacion.Deudas <= criterio.DeudaMaxima &&
			operacion.OperacionPagada == criterio.OperacionPagada {

			coincidencias++
		}
	}
	return coincidencias
}