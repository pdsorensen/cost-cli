package responses

import "github.com/pdsorensen/aws-cost-cli/cmd/models"

// ListJobRunsResponse ...
type ListJobRunsResponse struct {
	Runs    []models.DataBricksRun
	HasMore bool `json:"has_more"`
}
