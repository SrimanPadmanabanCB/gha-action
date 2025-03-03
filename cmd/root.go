package cmd

import (
	"context"
	"fmt"
	"github.com/calculi-corp/gha-register-build-artifact/internal/artifacts"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var (
	cmd = &cobra.Command{
		Use:   "gha-register-build-action",
		Short: "Publish the build artifact metadata to CloudBees Build Platform",
		Long:  "Publish the build artifact metadata to CloudBees Build Platform",
		RunE:  run,
	}
	cfg artifacts.Config
)

func Execute() error {
	return cmd.Execute()
}

func init() {
	err := setEnvVars(&cfg)
	if err != nil {
		panic(err)
	}
	setDefaultValues(&cfg)
}

func setDefaultValues(cfg *artifacts.Config) {
	artifactType := os.Getenv(artifacts.ArtifactType)
	if artifactType != "" {
		cfg.ArtifactType = artifactType
	} else {
		cfg.ArtifactType = ""
	}

	artifactDigest := os.Getenv(artifacts.ArtifactDigest)
	if artifactDigest != "" {
		cfg.ArtifactDigest = artifactDigest
	} else {
		cfg.ArtifactDigest = ""
	}
	cfg.ArtifactOperation = artifacts.PUBLISHED

}

func setEnvVars(cfg *artifacts.Config) error {
	ghaRunId := os.Getenv(artifacts.GithubRunId)
	if ghaRunId == "" {
		return fmt.Errorf(artifacts.GithubRunId + " is not set in the environment")
	}
	cfg.GhaRunId = ghaRunId

	ghaRunAttempt := os.Getenv(artifacts.GithubRunAttempt)
	if ghaRunAttempt == "" {
		return fmt.Errorf(artifacts.GithubRunAttempt + " is not set in the environment")
	}
	cfg.GhaRunAttempt = ghaRunAttempt

	cloudBeesApiUrl := os.Getenv(artifacts.CloudbeesApiUrl)
	if cloudBeesApiUrl == "" {
		return fmt.Errorf(artifacts.CloudbeesApiUrl + " is not set in the environment")
	}
	cfg.CloudBeesApiUrl = cloudBeesApiUrl

	cloudBeesApiToken := os.Getenv(artifacts.CloudbeesApiToken)
	if cloudBeesApiToken == "" {
		return fmt.Errorf(artifacts.CloudbeesApiToken + " is not set in the environment")
	}
	cfg.CloudBeesApiToken = cloudBeesApiToken

	artifactName := os.Getenv(artifacts.ArtifactName)
	if artifactName == "" {
		return fmt.Errorf(artifacts.ArtifactName + " is not set in the environment")
	}
	cfg.ArtifactName = artifactName

	artifactUrl := os.Getenv(artifacts.ArtifactUrl)
	if artifactUrl == "" {
		return fmt.Errorf(artifacts.ArtifactUrl + " is not set in the environment")
	}
	cfg.ArtifactUrl = artifactUrl

	artifactVersion := os.Getenv(artifacts.ArtifactVersion)
	if artifactVersion == "" {
		return fmt.Errorf(artifacts.ArtifactVersion + " is not set in the environment")
	}

	cfg.ArtifactVersion = artifactVersion

	repoFullName := os.Getenv(artifacts.RepoFullName)
	if repoFullName == "" {
		return fmt.Errorf(artifacts.RepoFullName + " is not set in the environment")
	}

	cfg.RepoFullName = repoFullName

	workflowName := os.Getenv(artifacts.WorkflowName)
	if workflowName == "" {
		return fmt.Errorf(artifacts.WorkflowName + " is not set in the environment")
	}

	cfg.WorkflowName = workflowName

	ghaRunNumber := os.Getenv(artifacts.GithubRunNumber)
	if ghaRunNumber == "" {
		return fmt.Errorf(artifacts.GithubRunNumber + " is not set in the environment")
	}

	cfg.GhaRunNumber = ghaRunNumber
	return nil
}

func run(_ *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("unknown arguments: %v", args)
	}
	newContext, cancel := context.WithCancel(context.Background())
	osChannel := make(chan os.Signal, 1)
	signal.Notify(osChannel, os.Interrupt)
	go func() {
		<-osChannel
		cancel()
	}()

	return cfg.Run(newContext)
}
