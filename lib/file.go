package lib

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

const FILENAME = "data/lazy.json"

func AddLazy(name, path string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	l := Lazy{name, path}
	index := data.findByName(name)

	if index != -1 {
		return errors.New(
			fmt.Sprintf(
				"A lazy named '%s' is already associated to '%s', to replace it use 'lazy replace %s %s'",
				name, data.Lazies[index].Path,
				name, data.Lazies[index].Path,
			),
		)
	}
	data.Lazies = append(data.Lazies, l)
	return writeData(data)
}

func RemoveLazy(name string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	index := data.findByName(name)

	if index == -1 {
		return errors.New(fmt.Sprintf("No lazy named '%s' found", name))
	}
	data.Lazies = append(data.Lazies[:index], data.Lazies[index+1:]...)
	return writeData(data)
}

func writeData(data *Lazies) error {
	b, err := json.Marshal(data)

	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(FILENAME, b, 0666); err != nil {
		return err
	}
	return nil
}

func getData() (*Lazies, error) {
	data := &Lazies{}
	f, err := ioutil.ReadFile(FILENAME)

	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(f, &data); err != nil {
		return nil, err
	}
	return data, nil
}
