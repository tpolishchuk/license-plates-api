package plate

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlateAPI struct {
	PlateService PlateService
}

func ProvidePlateAPI(p PlateService) PlateAPI {
	return PlateAPI{PlateService: p}
}

func (p *PlateAPI) FindAll(c *gin.Context) {
	plates := p.PlateService.FindAll()

	c.JSON(http.StatusOK, gin.H{"plates": ToPlateDTOs(plates)})
}

func (p *PlateAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	plate, err := p.PlateService.FindByID(uint(id))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plate": ToPlateDTO(plate)})
}

func (p *PlateAPI) Create(c *gin.Context) {
	var plateDTO PlateDTO
	err := c.BindJSON(&plateDTO)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdPlate := p.PlateService.Save(ToPlate(plateDTO))

	c.JSON(http.StatusCreated, gin.H{"plate": ToPlateDTO(createdPlate)})
}

func (p *PlateAPI) Update(c *gin.Context) {
	var plateDTO PlateDTO

	err := c.BindJSON(&plateDTO)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	plate, err := p.PlateService.FindByID(uint(id))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if plate == (Plate{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	plate.Abbreviation = plateDTO.Abbreviation
	plate.Country = plateDTO.Country
	plate.CityOrRegion = plateDTO.CityOrRegion
	plate.State = plateDTO.State

	p.PlateService.Save(plate)

	c.JSON(http.StatusOK, gin.H{"plate": ToPlateDTO(plate)})
}

func (p *PlateAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	plate, err := p.PlateService.FindByID(uint(id))

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if plate == (Plate{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.PlateService.Delete(plate)

	c.Status(http.StatusNoContent)
}
