package lib

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func BindNamePath(args []string) (string, string) {
	if len(args) != 2 {
		log.Fatal("Needs 2 arguments : name and path")
	}
	path, err := checkPath(args[1])

	if err != nil {
		log.Fatal(err)
	}
	return args[0], path
}

func BindNames(args []string) (string, string) {
	if len(args) != 2 {
		log.Fatal("Needs 2 arguments : name name")
	}
	return args[0], args[1]
}

func BindName(args []string) string {
	if len(args) != 1 {
		log.Fatal("Needs 1 argument : name")
	}
	return args[0]
}

func BindOptionalName(args []string) string {
	if len(args) != 1 {
		return ""
	}
	return args[0]
}

func checkPath(p string) (string, error) {
	path, _ := filepath.Abs(p)

	if _, err := os.Stat(path); err != nil {
		return "", errors.New(fmt.Sprintf("%s is not accessible", p))
	}
	return path, nil
}
