package models

type Cliente struct {
	ID               string      
	Sueldo           float64     
	Deudas           float64      
	OperacionHistorial []Operacion 
}
