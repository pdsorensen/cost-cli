package models


type DataBricksRun struct {
	JobID                int `json:"job_id"`
	RunID                int `json:"run_id"`
	NumberInJob          int `json:"number_in_job"`
	OriginalAttemptRunID int `json:"original_attempt_run_id"`
	State                struct {
		LifeCycleState string `json:"life_cycle_state"`
		ResultState    string `json:"result_state"`
		StateMessage   string `json:"state_message"`
	} `json:"state"`
	Task struct {
		NotebookTask struct {
			NotebookPath string `json:"notebook_path"`
		} `json:"notebook_task"`
	} `json:"task"`
	ClusterSpec struct {
		NewCluster struct {
			SparkVersion  string `json:"spark_version"`
			AwsAttributes struct {
				ZoneID              string `json:"zone_id"`
				FirstOnDemand       int    `json:"first_on_demand"`
				Availability        string `json:"availability"`
				InstanceProfileArn  string `json:"instance_profile_arn"`
				SpotBidPricePercent int    `json:"spot_bid_price_percent"`
				EbsVolumeCount      int    `json:"ebs_volume_count"`
			} `json:"aws_attributes"`
			NodeTypeID        string `json:"node_type_id"`
			EnableElasticDisk bool   `json:"enable_elastic_disk"`
			NumWorkers        int    `json:"num_workers"`
		} `json:"new_cluster"`
	} `json:"cluster_spec"`
	ClusterInstance struct {
		ClusterID      string `json:"cluster_id"`
		SparkContextID string `json:"spark_context_id"`
	} `json:"cluster_instance"`
	OverridingParameters struct {
		NotebookParams struct {
			P1Pids            string `json:"P1_Pids"`
			P2StartDate       string `json:"P2_StartDate"`
			P3EndDate         string `json:"P3_EndDate"`
			P4UseDummyData    string `json:"P4_UseDummyData"`
			P5OutputPath      string `json:"P5_OutputPath"`
			P6UseRestoredData string `json:"P6_UseRestoredData"`
		} `json:"notebook_params"`
	} `json:"overriding_parameters"`
	StartTime         int64  `json:"start_time"`
	SetupDuration     int    `json:"setup_duration"`
	ExecutionDuration int    `json:"execution_duration"`
	CleanupDuration   int    `json:"cleanup_duration"`
	Trigger           string `json:"trigger"`
	CreatorUserName   string `json:"creator_user_name"`
	RunName           string `json:"run_name"`
	RunPageURL        string `json:"run_page_url"`
	RunType           string `json:"run_type"`
}
