package plate

func ToPlate(plateDTO PlateDTO) Plate {
	return Plate{
		Abbreviation: plateDTO.Abbreviation,
		Country:      plateDTO.Country,
		CityOrRegion: plateDTO.CityOrRegion,
		State:        plateDTO.State}
}

func ToPlateDTO(plate Plate) PlateDTO {
	return PlateDTO{
		ID:           plate.ID,
		Abbreviation: plate.Abbreviation,
		Country:      plate.Country,
		CityOrRegion: plate.CityOrRegion,
		State:        plate.State}
}

func ToPlateDTOs(plates []Plate) []PlateDTO {
	plateDtos := make([]PlateDTO, len(plates))

	for i, itm := range plates {
		plateDtos[i] = ToPlateDTO(itm)
	}

	return plateDtos
}
