package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func writeData(data []Lazy) error {
	b, err := json.Marshal(data)

	if err != nil {
		return err
	}
	path := getJsonPath(FILENAME)

	if err := ioutil.WriteFile(path, b, 0666); err != nil {
		return err
	}
	return nil
}

func getData() ([]Lazy, error) {
	var data []Lazy
	path := getJsonPath(FILENAME)
	f, err := ioutil.ReadFile(path)

	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(f, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func getJsonPath(filename string) string {
	ex, err := os.Executable()

	if err != nil {
		panic(err)
	}
	return filepath.Join(filepath.Dir(ex), filename)
}
