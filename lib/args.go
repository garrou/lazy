package lib

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Bind2Args(args []string) (string, string) {
	if len(args) != 2 {
		log.Fatal("Needs 2 arguments : name and path")
	}
	err := checkPath(args[1])

	if err != nil {
		log.Fatal(err)
	}
	return args[0], args[1]
}

func BindArg(args []string) string {
	if len(args) != 1 {
		log.Fatal("Needs 1 argument : name")
	}
	return args[0]
}

func checkPath(p string) error {
	if _, err := os.Stat(p); err != nil {
		return errors.New(fmt.Sprintf("%s is not accessible", p))
	}
	return nil
}
