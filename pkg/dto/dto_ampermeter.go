package dto

import (
	"metrologARM/pkg/models"
	"time"
)

type AmpereMeterDTO struct {
	ID                     int                       `json:"id"`                     //id
	Title                  string                    `json:"title"`                  //наименвание
	Type                   string                    `json:"type"`                   //тип амперметра
	FactoryNumber          string                    `json:"factoryNumber"`          //заводской номер
	AccuracyClass          float32                   `json:"accuracyClass"`          //класс точности
	LimitMeasurement       string                    `json:"limitMeasurement"`       //предел измерение
	InstallationLocation   string                    `json:"installationLocation"`   //место установки
	Object                 models.Object             `json:"object"`                 //объект где установлен
	YearOfManufacture      int                       `json:"yearOfManufacture"`      //год выпуска
	DateOfVerification     time.Time                 `json:"dateOfVerification"`     //дата поверки или калибровки
	Periodicity            int                       `json:"periodicity"`            //периодичность
	DateOfLastVerification time.Time                 `json:"dateOfLastVerification"` //дата следующей поверки и калибровки
	TypesOfMeasurement     models.TypesOfMeasurement `json:"typesOfMeasurement"`
}

func NewAmpereMeterDTO(amper *models.AmpereMeter) *AmpereMeterDTO {
	return &AmpereMeterDTO{
		ID:                     amper.ID,
		Title:                  amper.Title,
		Type:                   amper.Type,
		FactoryNumber:          amper.FactoryNumber,
		AccuracyClass:          amper.AccuracyClass,
		LimitMeasurement:       amper.LimitMeasurement,
		InstallationLocation:   amper.InstallationLocation,
		Object:                 amper.Object,
		YearOfManufacture:      amper.YearOfManufacture,
		DateOfVerification:     amper.DateOfVerification,
		Periodicity:            amper.Periodicity,
		DateOfLastVerification: amper.DateOfLastVerification,
		TypesOfMeasurement:     amper.TypesOfMeasurement,
	}
}

func NewAmpereMeterListDTO(amperList []*models.AmpereMeter) []*AmpereMeterDTO {
	var amperDTOList []*AmpereMeterDTO
	for _, amper := range amperList {
		amperDTOList = append(amperDTOList, NewAmpereMeterDTO(amper))
	}
	return amperDTOList
}
