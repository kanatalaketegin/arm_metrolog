package models

/*
* модель для электросчетчика
 */

import "time"

type ElektricityMeter struct {
	ID                     int                `json:"id"`                     //id
	Title                  string             `json:"title"`                  //наименвание
	Type                   string             `json:"type"`                   //тип счетчика
	FactoryNumber          string             `json:"factoryNumber"`          //заводской номер
	AccuracyClass          float32            `json:"accuracyClass"`          //класс точности
	InstallationLocation   string             `json:"installationLocation"`   //место установки
	Object                 Object             `json:"object"`                 //объект где установлен
	DesignFactor           int                `json:"designFactor"`           //расчетный коэффициент
	DesignFactorTT         string             `json:"designFactorTT"`         //расчетный коэффициент ТТ
	DesignFactorTN         string             `json:"designFactorTN"`         //расчетный коэффициент ТН
	YearOfManufacture      time.Month         `json:"yearOfManufacture"`      //год выпуска
	DateOfVerification     time.Time          `json:"dateOfVerification"`     //дата поверки или калибровки
	Periodicity            int                `json:"periodicity"`            //периодичность
	DateOfLastVerification time.Time          `json:"dateOfLastVerification"` //дата следующей поверки и калибровки
	TypesOfMeasurement     TypesOfMeasurement `json:"typesOfMeasurement"`
}
