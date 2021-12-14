package models

type Student struct {
	id int `json:"ID"`
	Name string `json:"name"`
	Npm int `json:"NPM"`
	Address string `json:"address"`
	City string `json:"city"`
	Gander string `json:"gender"`
	Religion string `json:"religion"`
	Phone int `json:"phone"`
}