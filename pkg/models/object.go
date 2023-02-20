package models

/*
* модель для объекта
 */

type Object struct {
	ID      int     `json:"id"`      //id объекта
	Title   string  `json:"title"`   //короткое название
	Name    string  `json:"name"`    //название объекта
	Company Company `json:"company"` //на каком компании подлежить объект
}
