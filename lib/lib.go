package lib

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/cheynewallace/tabby"
)

const FILENAME = ".config/lazy.json"

func AddLazy(name, path string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	index := findByName(data, name)

	if index != -1 {
		return errors.New(
			fmt.Sprintf(
				"A lazy named '%s' is already associated to '%s', to replace it use 'lazy replace %s %s'",
				name, data[index].Path,
				name, data[index].Path,
			),
		)
	}
	data = append(data, Lazy{name, path})
	return writeData(data)
}

func GetLazies(name string) ([]Lazy, error) {
	data, err := getData()

	if err != nil {
		return nil, err
	}
	if name == "" {
		return data, nil
	}
	index := findByName(data, name)

	if index == -1 {
		return nil, errors.New(fmt.Sprintf("No lazy named '%s' found", name))
	}
	lazies := make([]Lazy, 0)
	return append(lazies, data[index]), nil
}

func RemoveLazy(name string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	index := findByName(data, name)

	if index == -1 {
		return errors.New(fmt.Sprintf("No lazy named '%s' found", name))
	}
	data = append(data[:index], data[index+1:]...)
	return writeData(data)
}

func ReplaceLazy(name, path string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	index := findByName(data, name)

	if index == -1 {
		return errors.New(fmt.Sprintf("No lazy named '%s' found", name))
	}
	data[index].Path = path
	return writeData(data)
}

func CopyLazy(oldName, newName string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	oldIndex := findByName(data, oldName)

	if oldIndex == -1 {
		return errors.New(fmt.Sprintf("No lazy named '%s' found", oldName))
	}
	newIndex := findByName(data, newName)

	if newIndex != -1 {
		return errors.New(fmt.Sprintf("A lazy named '%s' is already associated to '%s'", newName, data[newIndex].Path))
	}
	data = append(data, Lazy{newName, data[oldIndex].Path})
	return writeData(data)
}

func MoveLazy(oldName, newName string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	oldIndex := findByName(data, oldName)

	if oldIndex == -1 {
		return errors.New(fmt.Sprintf("No lazy named '%s' found", oldName))
	}
	data[oldIndex].Name = newName
	return writeData(data)
}

func Display(lazies []Lazy) {
	t := tabby.New()
	t.AddHeader("name", "path")

	if len(lazies) > 1 {
		for _, l := range lazies {
			t.AddLine(l.Name, l.Path)
		}
		t.Print()
		return
	}
	if len(lazies) > 0 {
		fmt.Println(lazies[0].Path)
	}
}

func writeData(data []Lazy) error {
	b, err := json.Marshal(data)

	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(FILENAME, b, 0666); err != nil {
		return err
	}
	return nil
}

func getData() ([]Lazy, error) {
	var data []Lazy
	f, err := ioutil.ReadFile(FILENAME)

	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(f, &data); err != nil {
		return nil, err
	}
	return data, nil
}
