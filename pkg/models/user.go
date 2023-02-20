package models

/*
* модель для пользователя
 */

type User struct {
	ID       int     `json:"id"`       //id
	Name     string  `json:"name"`     //имя пользователя
	LastName string  `json:"lastName"` //фамилия пользователя
	Position string  `json:"position"` //должность
	Company  Company `json:"company"`  //на каком компании работает
}
