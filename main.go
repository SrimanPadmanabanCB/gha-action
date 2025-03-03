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

	fmt.Printf("✅ Received message: %s\n", os.Args[2])

	runID := os.Getenv("GITHUB_RUN_ID")

	// If empty, provide fallback
	if runID == "" {
		runID = "Unknown Run ID"
	}

	fmt.Println("🚀 GitHub Run ID:", runID)

	runAttempt := os.Getenv("GITHUB_RUN_ATTEMPT")

	// If empty, provide fallback
	if runAttempt == "" {
		runAttempt = "Unknown Run Attempt"
	}

	fmt.Println("🚀 GitHub Run Attempt:", runAttempt)
	// Emit output so it can be consumed in GitHub Actions Workflow
	fmt.Printf("result=Processed: %s\n", inputMessage)
}
