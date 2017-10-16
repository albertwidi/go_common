package main

import (
	"github.com/albert-widi/go_common/logger"
)

func main() {
	l := logger.New()
	// l.AddTags("tag1", "tag2")
	l.WithFields(logger.Fields{"field1": "value1"}).Print("Haloha")
	l.Print("Haloha")
}
