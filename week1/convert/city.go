package convert

import "log"

type City struct {
	ID        string
	Name      string
	Districts map[string]District
}

func InsertDistrict(city City, districtSlice []string) error {
	districtName := districtSlice[0]
	districtID := districtSlice[1]
	districts := city.Districts
	if _, ok := districts[districtID]; !ok {
		district := District{
			ID:    districtID,
			Name:  districtName,
			Wards: make(map[string]Ward),
		}
		districts[districtID] = district
	} else {
		log.Println("District ", districtID, "existed")
	}
	return InsertWard(districts[districtID], districtSlice[2:])
}
