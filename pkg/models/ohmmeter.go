package models

import "time"

/*
* модель амперметра Омметры и магазин сопротивлений
 */

type Ohmmeter struct {
	ID                     int                `json:"id"`                     //id
	Title                  string             `json:"title"`                  //наименвание
	Type                   string             `json:"type"`                   //тип омметр или магазин сопративление
	FactoryNumber          string             `json:"factoryNumber"`          //заводской номер
	AccuracyClass          float32            `json:"accuracyClass"`          //класс точности
	LimitMeasurement       string             `json:"limitMeasurement"`       //предел измерение
	InstallationLocation   string             `json:"installationLocation"`   //место установки
	Object                 Object             `json:"object"`                 //объект где установлен
	YearOfManufacture      time.Month         `json:"yearOfManufacture"`      //год выпуска
	DateOfVerification     time.Time          `json:"dateOfVerification"`     //дата поверки или калибровки
	Periodicity            int                `json:"periodicity"`            //периодичность
	DateOfLastVerification time.Time          `json:"dateOfLastVerification"` //дата следующей поверки и калибровки
	TypesOfMeasurement     TypesOfMeasurement `json:"typesOfMeasurement"`
}
