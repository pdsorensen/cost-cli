/*
Copyright © 2020 Patrikk Dyrberg Sørensen <patri-kk@hotmail.com>

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
	"time"

	"github.com/pdsorensen/aws-cost-cli/cmd/responses"
	"github.com/pdsorensen/aws-cost-cli/cmd/services"
	"github.com/spf13/cobra"
)

// databricksCmd represents the databricks command
var databricksCmd = &cobra.Command{
	Use:   "databricks",
	Short: "Analyzes your databricks resources",
	Long: `
		              _,.
           ,''   '.     __....__ 
         ,'        >.-''        ''-.__,)
       ,'      _,''           _____ _,'
      /      ,'           _.:':::_':-._ 
     :     ,'       _..-''  \''.;.'-:::':. 
     ;    /       ,'  ,::'  .\,''.'. '\::)'  
    /    /      ,'        \   '. '  )  )/ 
   /    /      /:'.     '--''   \     ''
   '-._/      /::::)             )
      /      /,-.:(   , _   '.-' 
     ;      :(,'.'-' ','.     ;
    :       |:\'' )      '-.._\ _
    |         ':-(             ')''-._ 
    |           '.'.        /'''      '':-.-__,
    :           / ':\ .     :            ' \'-
     \        ,'   '}  '.   |
  _..-'.    ,''-.   }   |'-'    
,'__    '-'' -.'.'._|   | 
    '''--..,.__.(_|.|   |::._
      __..','/ ,' :  '-.|::)_'.
      '..__..-'   |'.      __,' 
                  :  '-._ '  ;
                   \ \   )  /
                   .\ '.   /
                    ::    /
                    :|  ,'
                    :;,
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var JobID int
var Token string
var DatabricksURL string

var listJobRuns = &cobra.Command{
	Use:   "jobs",
	Short: "Visualizes Databrick Job runs",
	Run: func(cmd *cobra.Command, args []string) {
		// 1: initialize services and variables
		service := services.NewDatabricksService(Token, DatabricksURL)
		chartService := services.NewChartService()

		// 2: Get results
		results := service.ListJobRuns(JobID)
		dates, executionTimes, spotBids := generateJobChartValues(results)
		fileLocation := chartService.GenerateJobChart(dates, executionTimes)

		// 3: Print outcomes
		for i := 0; i < len(executionTimes); i++ {
			fmt.Printf("date: %v, result: %v minutes, spotBid: %d \n", dates[i], executionTimes[i], spotBids[i])
		}

		fmt.Println("Generated file: ", fileLocation)
	},
}

func generateJobChartValues(results responses.ListJobRunsResponse) ([]string, []int, []int) {
	executionTimes := []int{}
	dates := []string{}
	spotBids := []int{}

	for _, item := range results.Runs {
		// 1: format ms to minutes
		startTime := time.Unix(0, item.StartTime*int64(time.Millisecond))
		converted := float64(item.ExecutionDuration)
		minutes := converted * 0.000016667

		// 2: append to result arrays
		dates = append(dates, startTime.Format("2006-01-02"))
		executionTimes = append(executionTimes, int(minutes))
		spotBids = append(spotBids, item.ClusterSpec.NewCluster.AwsAttributes.SpotBidPricePercent)
	}

	return dates, executionTimes, spotBids
}

func init() {
	listJobRuns.Flags().IntVarP(&JobID, "job", "j", 0, "Job ID that you will be analyzed")
	listJobRuns.Flags().StringVarP(&Token, "token", "t", "", "Bearer token")
	listJobRuns.Flags().StringVarP(&DatabricksURL, "url", "u", "", "BaseURL for your databricks services")

	listJobRuns.MarkFlagRequired("job")
	listJobRuns.MarkFlagRequired("token")
	listJobRuns.MarkFlagRequired("url")

	listJobRuns.AddCommand(analyzeCmd)
	databricksCmd.AddCommand(listJobRuns)
	rootCmd.AddCommand(databricksCmd)
}
