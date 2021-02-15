package humanity

import (
	"SoftwareGoDay1/Day1/data"
	"errors"
	"fmt"
	"strconv"
)

type Human struct {
	Name    string
	Age     int
	Country string
}

func NewHumanFromCSV(csv []string) (*Human, error) {
	newage, err := strconv.Atoi(csv[1])
	if err != nil {
		return nil, errors.New("age is not an int")
	}
	h := Human{csv[0], newage, csv[2]}
	return &h, nil
}

func NewHumansFromCsvFile(path string) ([]*Human, error) {
	list, err := data.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lenlist := len(list)
	h := make([]*Human, lenlist)
	for i, line := range list {
		lineCsv, err := data.LineToCSV(line)
		if err != nil {
			return nil, err
		}
		h[i], err = NewHumanFromCSV(lineCsv)
		if err != nil {
			return nil, err
		}
	}
	return h, err
}

func (h *Human) String() string {
	return fmt.Sprintf("Name : %s, Age : %d, Country : %s\n", h.Name, h.Age, h.Country)
}
