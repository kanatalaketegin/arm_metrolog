package models

/*
* модель для Компании
 */

type Company struct {
	ID    int    `json:"id"`    //id company
	Title string `json:"title"` //короткое название
	Name  string `json:"name"`  //name company
}
