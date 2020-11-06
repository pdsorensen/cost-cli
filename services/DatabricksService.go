package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pdsorensen/aws-cost-cli/cmd/models"
	"github.com/pdsorensen/aws-cost-cli/cmd/responses"
)

func getAccessToken() {}

// DataBricksService is the service
type DataBricksService struct {
	client  *http.Client
	token   string
	baseURL string
}

// ListJobRuns lists jobs from databricks
func (d DataBricksService) ListJobRuns(jobID int) responses.ListJobRunsResponse {
	endpoint := fmt.Sprintf("%s/api/2.0/jobs/runs/list?job_id=%v", d.baseURL, jobID)
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", d.token)

	res, err := d.client.Do(req)

	if err != nil {
		fmt.Println("Error occoured", err.Error())
	}

	decodedResponse := new(responses.ListJobRunsResponse)
	err = json.NewDecoder(res.Body).Decode(&decodedResponse)

	return *decodedResponse
}

// ListJobs ...
func (d DataBricksService) ListJobs() []models.DatabricksJob {
	// 1: Setup request
	endpoint := fmt.Sprintf("%s/api/2.0/jobs/list", d.baseURL)
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", d.token)

	// 2: Execute and validate errors
	res, err := d.client.Do(req)

	if err != nil {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println("Error occoured in request: ", err.Error(), body)
	}

	decodedResponse := new(responses.ListJobsResponse)
	err = json.NewDecoder(res.Body).Decode(&decodedResponse)

	if err != nil {
		fmt.Println(res.StatusCode)
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println("Error occoured when decoding: ", err.Error(), body)
	}

	return decodedResponse.Jobs
}

// NewDatabricksService ...
func NewDatabricksService(token string, baseURL string) DataBricksService {
	return DataBricksService{
		token:   "Bearer " + token,
		baseURL: baseURL,
		client:  &http.Client{},
	}
}
