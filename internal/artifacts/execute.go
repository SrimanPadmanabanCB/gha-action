package artifacts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

func (config *Config) Run(_ context.Context) (err error) {

	cloudEventData := prepareCloudEventData(config)

	cloudEvent, err := prepareCloudEvent(config, cloudEventData)
	if err != nil {
		return err
	}
	err = sendCloudEvent(cloudEvent, config)
	if err != nil {
		return err
	}
	return nil
}

func prepareCloudEvent(config *Config, output Output) (event.Event, error) {
	cloudEvent := cloudevents.NewEvent()
	cloudEvent.SetID(uuid.NewString())
	cloudEvent.SetSubject(config.WorkflowName + "|" + config.GhaRunId + "|" + config.GhaRunAttempt + "|" + config.GhaRunNumber)
	cloudEvent.SetType(BuildArtifactType)
	cloudEvent.SetSource(config.RepoFullName)
	cloudEvent.SetSpecVersion(SpecVersion)
	cloudEvent.SetTime(time.Now())
	err := cloudEvent.SetData(ContentTypeJson, output)
	if err != nil {
		return event.Event{}, fmt.Errorf("failed to set data: %v", err)
	}
	return cloudEvent, nil
}

func prepareCloudEventData(config *Config) Output {

	artifactInfo := &ArtifactInfo{
		ArtifactName:      config.ArtifactName,
		ArtifactUrl:       config.ArtifactUrl,
		ArtifactVersion:   config.ArtifactVersion,
		ArtifactType:      config.ArtifactType,
		ArtifactDigest:    config.ArtifactDigest,
		ArtifactOperation: config.ArtifactOperation,
		RunId:             config.GhaRunId,
		RunAttempt:        config.GhaRunAttempt,
	}
	output := Output{
		ArtifactInfo: *artifactInfo,
	}
	return output
}
func sendCloudEvent(cloudEvent event.Event, config *Config) error {
	eventJSON, err := json.Marshal(cloudEvent)
	if err != nil {
		return fmt.Errorf("Error encoding CloudEvent JSON:", err)

	}
	req, _ := http.NewRequest(PostMethod, config.CloudBeesApiUrl, bytes.NewBuffer(eventJSON))

	// For Local Testing
	//req, _ := http.NewRequest(PostMethod, "http://localhost:8080/events", bytes.NewBuffer(eventJSON))

	req.Header.Set(ContentTypeHeaderKey, ContentTypeCloudEventsJson)
	req.Header.Set(AuthorizationHeaderKey, Bearer+config.CloudBeesApiToken)
	client := &http.Client{}
	resp, err := client.Do(req) // Fire and forget
	if err != nil {
		return fmt.Errorf("error sending CloudEvent", err)

	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	fmt.Println("CloudEvent sent successfully!")
	return nil
}
