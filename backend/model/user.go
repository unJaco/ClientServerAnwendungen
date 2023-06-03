package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  		string 		`json:"first_name"`
	LastName  		string 		`json:"last_name"`
	TelNr  			string 		`json:"tel_nr"`
	PwHash  		string 		`json:"pw_hash"`
	// grom:"unique" sorgt dafür, dass das Feld email unique sein muss
	Email 			string	 	`json:"email" gorm:"unique"`
}	



