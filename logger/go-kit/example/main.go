package main

import (
	stderrs "errors"

	"github.com/albert-widi/go_common/errors"
	"github.com/albert-widi/go_common/logger"
)

func main() {
	l := logger.New()
	// l.AddTags("tag1", "tag2")
	// l.WithFields(logger.Fields{"field1": "value1"}).Print("Haloha")
	// l.Print("Haloha")

	err := errors.New("Ini error aslinya", errors.Fields{"map1": "value1"})
	err2 := errors.New("Ini error aslinya", errors.Fields{"map2": "value2"})
	err3 := stderrs.New("Ini error aslinya")
	if errors.Match(err3, err2) {
		l.Errors(err)
	} else {
		l.Print("Not match")
	}
}
