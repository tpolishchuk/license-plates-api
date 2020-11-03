package plate

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PlateRepository struct {
	DB *gorm.DB
}

func ProvidePlateRepostiory(DB *gorm.DB) PlateRepository {
	return PlateRepository{DB: DB}
}

func (p *PlateRepository) FindAll() []Plate {
	var plates []Plate
	p.DB.Find(&plates)

	return plates
}

func (p *PlateRepository) FindByID(id uint) (Plate, error) {
	var plate Plate
	p.DB.First(&plate, id)

	if plate.ID == 0 {
		return Plate{}, errors.New("Plate cannot be found by a given id")
	}

	return plate, nil
}

func (p *PlateRepository) Save(plate Plate) Plate {
	p.DB.Save(&plate)

	return plate
}

func (p *PlateRepository) Delete(plate Plate) {
	p.DB.Delete(&plate)
}
