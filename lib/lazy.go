package lib

import (
	"fmt"
	"strings"
)

type Lazy struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func findByName(lazies []Lazy, name string) int {
	for i, l := range lazies {
		if strings.EqualFold(name, l.Name) {
			return i
		}
	}
	return -1
}

func (l Lazy) String() string {
	return fmt.Sprintf("%s : %s", l.Name, l.Path)
}
