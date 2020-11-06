/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/pdsorensen/aws-cost-cli/cmd/services"
	"github.com/spf13/cobra"
)

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "A brief description of your command",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1: Get select options
		service := services.NewDatabricksService(Token, DatabricksURL)
		jobs := service.ListJobs()
		names := []string{}

		for _, job := range jobs {
			names = append(names, job.Settings.Name)
		}

		// 2: Handle and setup promt
		prompt := promptui.Select{
			Label: "Job",
			Items: names,
		}

		index, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Something went wrong %v\n", err)
			return
		}

		fmt.Printf("You choose %v (%d) (id=%d)\n", result, index, jobs[index].JobID)
	},
}

func init() {
	analyzeCmd.Flags().StringVarP(&Token, "token", "t", "", "Bearer token")
	analyzeCmd.Flags().StringVarP(&DatabricksURL, "url", "u", "", "BaseURL for your databricks services")
}
