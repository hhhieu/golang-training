package convert

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// ConvertCSV2YML returns nil if success and not nil if failed.
func ConvertCSV2YML(csvFile string) error {
	f, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer f.Close() // this needs to be after the err check
	// 	r := strings.NewReader(`Tỉnh Thành Phố,Mã TP,Quận Huyện,Mã QH,Phường Xã,Mã PX,Cấp,Tên Tiếng Anh
	// Thành phố Hà Nội,01,Quận Ba Đình,001,Phường Phúc Xá,00001,Phường,
	// Thành phố Hà Nội,01,Huyện Đông Anh,017,Xã Tầm Xá,00517,Xã,
	// Tỉnh Thanh Hóa,38,Huyện Vĩnh Lộc,393,Xã Vĩnh Long,15361,Xã,`)
	// csvReader := csv.NewReader(r)
	// csvReader.FieldsPerRecord = -1
	// records, err := csvReader.ReadAll()
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	cities := make(map[string]City)
	for i, record := range records {
		// Skip header
		if i < 1 {
			continue
		}
		insertCity(cities, record)
		fmt.Println(i, " | ", record)
	}
	d, err := yaml.Marshal(&cities)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("YAML")
	fmt.Println(string(d))
	err = ioutil.WriteFile("out.yml", d, 0o644)
	return nil
}
