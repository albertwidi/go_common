package main

import (
	"github.com/albert-widi/go_common/logger"
)

func main() {
	l := logger.New()
	l.WithFields(logger.Fields{"field1": "value1"}).Print("Haloha")
	l.Print("Haloha")
}
