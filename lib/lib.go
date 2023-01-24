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
	data.Lazies = append(data.Lazies, Lazy{name, path})
	return writeData(data)
}

func GetLazies(name string) ([]Lazy, error) {
	data, err := getData()

	if err != nil {
		return nil, err
	}
	if name == "" {
		return data.Lazies, nil
	}
	index := data.findByName(name)

	if index == -1 {
		return nil, errors.New(fmt.Sprintf("No lazy named '%s' found", name))
	}
	lazies := make([]Lazy, 0)
	return append(lazies, data.Lazies[index]), nil
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

func ReplaceLazy(name, path string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	index := data.findByName(name)

	if index == -1 {
		return errors.New(fmt.Sprintf("No lazy named '%s' found", name))
	}
	data.Lazies[index].Path = path
	return writeData(data)
}

func CopyLazy(oldName, newName string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	oldIndex := data.findByName(oldName)

	if oldIndex == -1 {
		return errors.New(fmt.Sprintf("No lazy named '%s' found", oldName))
	}
	newIndex := data.findByName(newName)

	if newIndex != -1 {
		return errors.New(fmt.Sprintf("A lazy named '%s' is already associated to '%s'", newName, data.Lazies[newIndex].Path))
	}
	data.Lazies = append(data.Lazies, Lazy{newName, data.Lazies[oldIndex].Path})
	return writeData(data)
}

func MoveLazy(oldName, newName string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	oldIndex := data.findByName(oldName)

	if oldIndex == -1 {
		return errors.New(fmt.Sprintf("No lazy named '%s' found", oldName))
	}
	data.Lazies[oldIndex].Name = newName
	return writeData(data)
}

func Display(lazies []Lazy) {
	for _, c := range lazies {
		fmt.Println(c)
	}
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
