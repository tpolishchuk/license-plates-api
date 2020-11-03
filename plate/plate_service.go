package plate

type PlateService struct {
	PlateRepository PlateRepository
}

func ProvidePlateService(p PlateRepository) PlateService {
	return PlateService{PlateRepository: p}
}

func (p *PlateService) FindAll() []Plate {
	return p.PlateRepository.FindAll()
}

func (p *PlateService) FindByID(id uint) (Plate, error) {
	return p.PlateRepository.FindByID(id)
}

func (p *PlateService) Save(plate Plate) Plate {
	p.PlateRepository.Save(plate)

	return plate
}

func (p *PlateService) Delete(plate Plate) {
	p.PlateRepository.Delete(plate)
}
