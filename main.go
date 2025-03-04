package main

import (
	"fmt"
	"gha-register-build-artifact/cmd"
	"log"
	"os"
	"strings"
)

func main() {

	for _, env := range os.Environ() {
		// Split into key and value
		pair := strings.SplitN(env, "=", 2)
		fmt.Printf("🔹 %s = %s\n", pair[0], pair[1])
	}
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
