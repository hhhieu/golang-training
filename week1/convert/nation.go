package convert

import "log"

func insertCity(nation map[string]City, nationSlice []string) error {
	cityName := nationSlice[0]
	cityID := nationSlice[1]
	if _, ok := nation[cityID]; !ok {
		city := City{
			ID:        cityID,
			Name:      cityName,
			Districts: make(map[string]District),
		}
		nation[cityID] = city
	} else {
		log.Println("City ", cityID, "existed")
	}
	return InsertDistrict(nation[cityID], nationSlice[2:])
}
