package main

import (
	"fmt"
	"math"
)

type city struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Visited   bool    `json:"visited"`
}

type route struct {
	Cities   []city  `json:"cities"`
	Distance float64 `json:"distance"`
}

type citiesRepository interface {
	getAll() ([]city, error)
}

type previousCitiesRepository struct {
	cities []city
}

func (repo *previousCitiesRepository) getAll() ([]city, error) {
	return repo.cities, nil
}

type findRouteUserCase interface {
	execute(cities []city) (route, error)
}

type nearestInsertionHeuristic struct {
	repository citiesRepository
}

func (n *nearestInsertionHeuristic) execute(cities []city) (route, error) {
	if len(cities) < 2 {
		return route{}, fmt.Errorf("minimum two cities required")
	}

	for i := range cities {
		cities[i].Visited = false
	}

	currentCity := cities[0]
	currentCity.Visited = true
	remainingCities := cities[1:]

	route := route{Cities: []city{currentCity}}
	for len(remainingCities) > 0 {
		var nearestCityIndex int
		nearestDistance := math.MaxFloat64

		for i, c := range remainingCities {
			if !c.Visited && distance(currentCity, c) < nearestDistance {
				nearestCityIndex = i
				nearestDistance = distance(currentCity, c)
			}
		}
		currentCity = remainingCities[nearestCityIndex]
		currentCity.Visited = true
		route.Cities = append(route.Cities, currentCity)
		route.Distance += nearestDistance

		remainingCities = append(remainingCities[:nearestCityIndex], remainingCities[nearestCityIndex+1:]...)
	}
	route.Distance += distance(route.Cities[len(route.Cities)-1], route.Cities[0])

	return route, nil
}

func distance(firstCity city, secoundCity city) float64 {
	return math.Sqrt(math.Pow(firstCity.Longitude-secoundCity.Longitude, 2) + math.Pow(firstCity.Latitude-secoundCity.Latitude, 2))
}

func main() {
	cities := []city{
		{Name: "Mata Grande", Latitude: -9.118243, Longitude: -37.732298},
		{Name: "Presidente Figueiredo", Latitude: -2.029813, Longitude: -60.02338},
		{Name: "Gramado dos Loureiros", Latitude: -27.442895, Longitude: -52.91486},
		{Name: "Casa Nova", Latitude: -9.164083, Longitude: -40.973977},
		{Name: "Pentecoste", Latitude: -3.792736, Longitude: -39.269159},
		{Name: "Gaspar", Latitude: -26.933597, Longitude: -48.953428},
		{Name: "Irupi", Latitude: -20.350122, Longitude: -41.644359},
		{Name: "Campos Belos", Latitude: -13.034984, Longitude: -46.768112},
		{Name: "Indiara", Latitude: -17.138739, Longitude: -49.986162},
		{Name: "Cândido Mendes", Latitude: -1.43265, Longitude: -45.716112},
	}

	repository := previousCitiesRepository{cities: cities}
	userCase := nearestInsertionHeuristic{repository: &repository}

	route, err := userCase.execute(cities)
	if err != nil {
		fmt.Println("error finding route:", err)
		return
	}

	fmt.Println("route:", route)
	for _, city := range route.Cities {
		fmt.Println("-", city.Name)
	}
	fmt.Println("total distance:", route.Distance)
}
