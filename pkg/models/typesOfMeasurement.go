package models

/*
* модель поверка или калибровка
 */

const (
	verification = "поверка"
	calibration  = "калибровка"
)

type TypesOfMeasurement struct {
	ID              int    `json:"ID"`
	NameMeasurement string `json:"nameMeasurement"`
}
