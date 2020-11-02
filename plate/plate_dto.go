package plate

type PlateDTO struct {
	ID           uint   `json:"id,string,omitempty"`
	Abbreviation string `json:"abbreviation"`
	Country      string `json:"country"`
	CityOrRegion string `json:"city_or_region"`
	State        string `json:"state"`
}
