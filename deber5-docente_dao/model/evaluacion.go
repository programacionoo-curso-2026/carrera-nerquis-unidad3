package model

// Evaluacion representa la evaluación de un docente.
type Evaluacion struct {
	ID            string
	DocenteID     string
	CompetenciaID string
	Calificacion  float64
	Observacion   string
}
