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
	data, err := getLazies()

	if err != nil {
		return err
	}
	l := Lazy{name, path}
	dup := data.findByName(name)

	if dup != -1 {
		return errors.New(
			fmt.Sprintf(
				"A lazy called '%s' is already associated to '%s', to replace it use 'lazy replace %s %s'",
				name, data.Lazies[dup].Path,
				name, data.Lazies[dup].Path,
			),
		)
	}
	data.Lazies = append(data.Lazies, l)
	b, mErr := json.Marshal(data)

	if mErr != nil {
		return mErr
	}
	if err := ioutil.WriteFile(FILENAME, b, 0666); err != nil {
		return err
	}
	return nil
}

func getLazies() (*Lazies, error) {
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
