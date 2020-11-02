package main

import (
	"license-plates-api/plate"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initPlateAPI(db *gorm.DB) plate.PlateAPI {
	wire.Build(plate.ProvidePlateRepostiory, plate.ProvidePlateService, plate.ProvidePlateAPI)

	return plate.PlateAPI{}
}
