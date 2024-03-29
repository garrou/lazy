package lib

import (
	"errors"
	"fmt"

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
				"error : a lazy named '%s' is already associated to '%s', to replace it use 'lazy replace %s [path]'",
				name,
				data[index].Path,
				name,
			),
		)
	}
	data = append(data, Lazy{name, path})
	return writeData(data)
}

func GetLazy(name string) (*Lazy, error) {
	data, err := getData()

	if err != nil {
		return nil, err
	}
	index := findByName(data, name)

	if index == -1 {
		return nil, errors.New(fmt.Sprintf("error : no lazy named '%s' found", name))
	}
	return &data[index], nil
}

func GetLazies() ([]Lazy, error) {
	data, err := getData()

	if err != nil {
		return nil, err
	}
	return data, nil
}

func RemoveLazy(name string) error {
	data, err := getData()

	if err != nil {
		return err
	}
	index := findByName(data, name)

	if index == -1 {
		return errors.New(fmt.Sprintf("error : no lazy named '%s' found", name))
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
		return errors.New(fmt.Sprintf("error : no lazy named '%s' found", name))
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
		return errors.New(fmt.Sprintf("error : no lazy named '%s' found", oldName))
	}
	newIndex := findByName(data, newName)

	if newIndex != -1 {
		return errors.New(fmt.Sprintf("error : a lazy named '%s' is already associated to '%s'", newName, data[newIndex].Path))
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
		return errors.New(fmt.Sprintf("error : no lazy named '%s' found", oldName))
	}
	data[oldIndex].Name = newName
	return writeData(data)
}

func Display(lazies []Lazy) {
	t := tabby.New()
	t.AddHeader("name", "path")

	for _, l := range lazies {
		t.AddLine(l.Name, l.Path)
	}
	t.Print()
}
