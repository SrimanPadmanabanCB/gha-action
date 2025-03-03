package artifacts

const (
	ArtifactName     = "ARTIFACT_NAME"
	ArtifactUrl      = "ARTIFACT_URL"
	ArtifactVersion  = "ARTIFACT_VERSION"
	ArtifactType     = "ARTIFACT_TYPE"
	ArtifactDigest   = "ARTIFACT_DIGEST"
	GithubRunId      = "GITHUB_RUN_ID"
	GithubRunAttempt = "GITHUB_RUN_ATTEMPT"
	GithubRunNumber  = "GITHUB_RUN_NUMBER"

	CloudbeesApiUrl   = "CLOUDBEES_API_URL"
	CloudbeesApiToken = "CLOUDBEES_API_TOKEN"
	PUBLISHED         = "PUBLISHED"
	RepoFullName      = "REPO_FULL_NAME"
	WorkflowName      = "WORKFLOW_NAME"

	BuildArtifactType          = "cloudbees.platform.register.build.artifact"
	SpecVersion                = "1.0"
	ContentTypeJson            = "application/json"
	ContentTypeHeaderKey       = "Content-Type"
	ContentTypeCloudEventsJson = "application/cloudevents+json"
	AuthorizationHeaderKey     = "Authorization"
	Bearer                     = "Bearer "
	PostMethod                 = "POST"
)
