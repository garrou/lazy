package lib

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckArgsNamePath(args []string) (string, string, error) {
	if len(args) != 2 || strings.TrimSpace(args[0]) == "" || strings.TrimSpace(args[1]) == "" {
		return "", "", errors.New("needs 2 args, -h for help")
	}
	path, err := checkPath(args[1])

	if err != nil {
		return "", "", err
	}
	if strings.TrimSpace(args[0]) == "" {
		return "", "", errors.New("error : needs 2 args, -h for help")
	}
	return args[0], path, nil
}

func CheckArgs(args []string, expected int) error {
	if len(args) != expected {
		return errors.New(fmt.Sprintf("error : needs %d args, -h for help", expected))
	}
	for _, v := range args {
		if strings.TrimSpace(v) == "" {
			return errors.New(fmt.Sprintf("error : needs %d args, -h for help", expected))
		}
	}
	return nil
}

func checkPath(p string) (string, error) {
	path, _ := filepath.Abs(p)

	if _, err := os.Stat(path); err != nil {
		return "", errors.New(fmt.Sprintf("error : %s is not accessible", p))
	}
	return path, nil
}
