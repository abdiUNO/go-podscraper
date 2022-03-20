package main

import (
	log "github.com/sirupsen/logrus"
	"podscraper/cmd"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:    false,
		DisableTimestamp: false,
		PadLevelText:     true,
	})

	cmd.Execute()
}
