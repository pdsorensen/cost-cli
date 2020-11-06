package models

// DatabricksJob ...
type DatabricksJob struct {
	JobID    int `json:"job_id"`
	Settings struct {
		Name string `json:"name"`
		// NewCluster struct {
		// 	SparkVersion  string `json:"spark_version"`
		// 	AwsAttributes struct {
		// 		ZoneID              string `json:"zone_id"`
		// 		FirstOnDemand       int    `json:"first_on_demand"`
		// 		Availability        string `json:"availability"`
		// 		InstanceProfileArn  string `json:"instance_profile_arn"`
		// 		SpotBidPricePercent int    `json:"spot_bid_price_percent"`
		// 		EbsVolumeType       string `json:"ebs_volume_type"`
		// 		EbsVolumeCount      int    `json:"ebs_volume_count"`
		// 		EbsVolumeSize       int    `json:"ebs_volume_size"`
		// 	} `json:"aws_attributes"`
		// 	NodeTypeID   string `json:"node_type_id"`
		// 	SparkEnvVars struct {
		// 		PYSPARKPYTHON string `json:"PYSPARK_PYTHON"`
		// 	} `json:"spark_env_vars"`
		// 	EnableElasticDisk bool `json:"enable_elastic_disk"`
		// 	NumWorkers        int  `json:"num_workers"`
		// } `json:"new_cluster"`
		EmailNotifications struct {
			OnFailure             []string `json:"on_failure"`
			AlertOnLastAttempt    bool     `json:"alert_on_last_attempt"`
			NoAlertForSkippedRuns bool     `json:"no_alert_for_skipped_runs"`
		} `json:"email_notifications"`
		TimeoutSeconds int `json:"timeout_seconds"`
		Schedule       struct {
			QuartzCronExpression string `json:"quartz_cron_expression"`
			TimezoneID           string `json:"timezone_id"`
			PauseStatus          string `json:"pause_status"`
		} `json:"schedule"`
		NotebookTask struct {
			NotebookPath      string `json:"notebook_path"`
			RevisionTimestamp int    `json:"revision_timestamp"`
		} `json:"notebook_task"`
		MaxConcurrentRuns int `json:"max_concurrent_runs"`
	} `json:"settings"`
	CreatedTime     int64  `json:"created_time"`
	CreatorUserName string `json:"creator_user_name"`
}
