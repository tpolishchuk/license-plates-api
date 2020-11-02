package plate

import "github.com/jinzhu/gorm"

type Plate struct {
	gorm.Model
	Abbreviation string
	Country      string
	CityOrRegion string
	State        string
}
