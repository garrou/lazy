package lib

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func BindNamePath(args []string) (string, string, error) {
	if len(args) != 2 {
		return "", "", errors.New("needs 2 args, -h for help")
	}
	path, err := checkPath(args[1])

	if err != nil || strings.TrimSpace(args[0]) == "" {
		return "", "", errors.New("needs 2 args, -h for help")
	}
	return args[0], path, nil
}

func BindNames(args []string) (string, string, error) {
	if len(args) != 2 {
		return "", "", errors.New("needs 2 args, -h for help")
	}
	return args[0], args[1], nil
}

func BindName(args []string) (string, error) {
	if len(args) != 1 || strings.TrimSpace(args[0]) == "" {
		return "", errors.New("needs 1 arg, -h for help")
	}
	return args[0], nil
}

func checkPath(p string) (string, error) {
	path, _ := filepath.Abs(p)

	if _, err := os.Stat(path); err != nil {
		return "", errors.New(fmt.Sprintf("%s is not accessible", p))
	}
	return path, nil
}
