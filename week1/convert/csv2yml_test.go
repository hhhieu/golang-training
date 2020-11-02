package convert

import "testing"

func TestConvertCSV2YML_HappyCase(t *testing.T) {
	csvFile := "../data.csv"
	err := ConvertCSV2YML(csvFile)
	if err != nil {
		t.Error("This should not error", err)
	}
}

func TestConvertCSV2YML_FileNotFound(t *testing.T) {
	csvFile := "NonExist.csv"
	err := ConvertCSV2YML(csvFile)
	if err == nil {
		t.Error("This should not nil")
	}
}

func TestConvertCSV2YML_FileInvalid(t *testing.T) {
	csvFile := "../googlelogo.png"
	err := ConvertCSV2YML(csvFile)
	if err == nil {
		t.Error("This should not nil", err)
	}
}
