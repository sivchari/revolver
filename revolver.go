package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var fire bool
	revolver := os.Getenv("REVOLVER_BULLETS")
	arg := os.Args[1]
	bullets := strings.Split(revolver, ",")
	for _, bullet := range bullets {
		if arg == bullet {
			fire = true
			break
		}
	}
	if !fire {
		return
	}
	trigger := os.Getenv("REVOLVER_TRIGGER")
	log.Println(trigger)
	revolvers := strings.Split(trigger, " ")
	exec.Command(revolvers[0], revolvers[1:]...)
}
