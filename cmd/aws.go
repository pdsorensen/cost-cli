/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/pdsorensen/aws-cost-cli/cmd/services"
	"github.com/spf13/cobra"
)

var Profile string
var Region string

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("resource must be specified")
		}

		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(Region), config.WithSharedConfigProfile(Profile))
		if err != nil {
			fmt.Println(err.Error())
		}

		anomalies := services.GetAnomalies(cfg)

		for _, anomaly := range anomalies {
			fmt.Printf("[%s][%s] impact: %f$ \n", *anomaly.AnomalyStartDate, *anomaly.RootCauses[0].Service, anomaly.Impact.TotalImpact)
		}
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// awsCmd.PersistentFlags().String("foo", "", "A help for foo")
	awsCmd.Flags().StringVarP(&Profile, "profile", "p", "", "AWS Profile to use")
	awsCmd.Flags().StringVarP(&Region, "region", "r", "us-east-1", "AWS Region to use")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// awsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
