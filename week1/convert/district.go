package convert

import "errors"

type District struct {
	ID    string
	Name  string
	Wards map[string]Ward
}

type Ward struct {
	ID          string
	Name        string
	Type        string
	EnglishName string
}

func InsertWard(district District, wardSlice []string) error {
	wardName := wardSlice[0]
	wardID := wardSlice[1]
	wardType := wardSlice[2]
	wardEnglishName := wardSlice[3]
	wards := district.Wards
	if _, ok := wards[wardID]; !ok {
		ward := Ward{
			ID:          wardID,
			Name:        wardName,
			Type:        wardType,
			EnglishName: wardEnglishName,
		}
		wards[wardID] = ward
	} else {
		return errors.New("Ward existed")
	}
	return nil
}
