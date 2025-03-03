package main

import (
	"github.com/calculi-corp/gha-register-build-artifact/cmd"
	"log"
)

func main() {

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
