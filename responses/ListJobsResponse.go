package responses

import "github.com/pdsorensen/aws-cost-cli/cmd/models"

// ListJobsResponse ...
type ListJobsResponse struct {
	Jobs []models.DatabricksJob `json:"jobs"`
}
