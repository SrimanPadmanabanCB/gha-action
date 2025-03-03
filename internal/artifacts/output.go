package artifacts

type ArtifactInfo struct {
	ArtifactName      string `json:"artifactName,omitempty"`
	ArtifactUrl       string `json:"artifactUrl,omitempty"`
	ArtifactVersion   string `json:"artifactVersion,omitempty"`
	ArtifactType      string `json:"artifactType,omitempty"`
	ArtifactDigest    string `json:"artifactDigest,omitempty"`
	ArtifactOperation string `json:"artifactOperation,omitempty"`
	RunId             string `json:"runId,omitempty"`
	RunAttempt        string `json:"runAttempt,omitempty"`
}

type Output struct {
	ArtifactInfo ArtifactInfo `json:"artifactInfo,omitempty"`
}
