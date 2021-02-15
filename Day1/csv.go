package data

import (
	"errors"
	"io/ioutil"
	"strings"
)

func ReadFile(path string) ([]string, error) {
	str, err := ioutil.ReadFile("example.csv")
	if err != nil {
		return nil, err
	}
	list := strings.Split(string(str), "\n")
	return list, err
}

func LineToCSV(line string) ([]string, error) {
	list := strings.Split(line, ",")
	if len(list) != 3 {
		return nil, errors.New("Error list format")
	}
	return list, nil
}
