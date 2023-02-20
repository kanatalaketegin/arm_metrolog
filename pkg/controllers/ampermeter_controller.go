package controllers

import (
	"encoding/json"
	"log"
	"metrologARM/pkg/configs"
	"metrologARM/pkg/dto"
	"metrologARM/pkg/models"
	"net/http"
)

func CreateAmpereMeter(w http.ResponseWriter, r *http.Request) {
	var ampereMeterDto dto.AmpereMeterDTO
	err := json.NewDecoder(r.Body).Decode(&ampereMeterDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ampereMeter := models.AmpereMeter{
		Title:                  ampereMeterDto.Title,
		Type:                   ampereMeterDto.Type,
		FactoryNumber:          ampereMeterDto.FactoryNumber,
		AccuracyClass:          ampereMeterDto.AccuracyClass,
		LimitMeasurement:       ampereMeterDto.LimitMeasurement,
		InstallationLocation:   ampereMeterDto.InstallationLocation,
		Object:                 ampereMeterDto.Object,
		YearOfManufacture:      ampereMeterDto.YearOfManufacture,
		DateOfVerification:     ampereMeterDto.DateOfLastVerification,
		Periodicity:            ampereMeterDto.Periodicity,
		DateOfLastVerification: ampereMeterDto.DateOfLastVerification,
		TypesOfMeasurement:     ampereMeterDto.TypesOfMeasurement,
	}

	err = configs.DB.Create(&ampereMeter).Error
	if err != nil {
		log.Println(err)
		http.Error(w, "Error creating amperemeter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ampereMeter)
}
