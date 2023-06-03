package model

import (
	enums "github.com/unJaco/UberClientServer/backend/model/enums"
	"gorm.io/gorm"
)


type Vehicle struct {

	gorm.Model
	Plate 		string 						`json:"plate" gorm:"varchar(12);unique"`
	// ein driver kann nur ein fahrzeug haben
	DriverId    uint 						`json:"driverId" gorm:"unique"`
	// User muss vorhanden sein, damit GORM das feld "driverID" zu einem foreignKey machen kann
	User 		User						`json:"-" gorm:"foreignkey:DriverId"`
	// gorm:"embedded" sorgt dafür, dass die felder von location, also lat und lon, automatisch auch in der DB erscheinen
	Location    Location					`json:"location" gorm:"embedded"`
	Status 		enums.VehhicleStatus			`json:"status"`
	
}