package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if input was provided
	if len(os.Args) < 2 {
		fmt.Println("❌ Error: No input received. Exiting...")
		os.Exit(1)
	}

	// Capture input argument (GitHub Action input)
	inputMessage := os.Args[1]

	// Process logic (custom action behavior)
	fmt.Println("🔥 Running Custom GitHub Action in Go!")
	fmt.Printf("✅ Received message: %s\n", inputMessage)

	// Emit output so it can be consumed in GitHub Actions Workflow
	fmt.Printf("result=Processed: %s\n", inputMessage)
}