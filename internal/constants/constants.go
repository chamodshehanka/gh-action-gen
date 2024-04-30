package constants

const (
	ActionTypeAppCI      = "Continuous Integration (CI)"
	ActionTypeDeployment = "Continuous Deployment (CD)"
	ActionTypePublish    = "Publish Package"

	TriggerOnCommit           = "OnCommit"
	TriggerOnWorkflowDispatch = "OnWorkflowDispatch"
	TriggerOnPullRequest      = "OnPullRequest"
)
