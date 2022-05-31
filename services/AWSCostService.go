package services

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func GetAnomalies(cfg aws.Config) []types.Anomaly {
	now := time.Now()
	start := now.AddDate(0, -1, 0).Format("2006-01-02")
	end := now.Format("2006-01-02")

	interval := types.AnomalyDateInterval{
		StartDate: &start,
		EndDate:   &end,
	}

	svc := costexplorer.NewFromConfig(cfg)

	input := costexplorer.GetAnomaliesInput{
		DateInterval: &interval,
	}

	results, err := svc.GetAnomalies(context.TODO(), &input)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return results.Anomalies
}
