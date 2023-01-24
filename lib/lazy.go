package lib

import (
	"fmt"
	"strings"
)

type Lazies struct {
	Lazies []Lazy `json:"lazies"`
}

func (l Lazies) findByName(name string) int {
	for i, l := range l.Lazies {
		if strings.EqualFold(name, l.Name) {
			return i
		}
	}
	return -1
}

type Lazy struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (l Lazy) String() string {
	return fmt.Sprintf("%s : %s", l.Name, l.Path)
}
