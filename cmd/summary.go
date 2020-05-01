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
	"github.com/spf13/viper"
	"k8s-enhancements/sheets"
	"k8s-enhancements/utils"

	"github.com/spf13/cobra"
)

var assignee string

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Get Summary of All the items being Tracked",
	Long: `Get Summary of All the items being Tracked`,
	PreRun: func(cmd *cobra.Command, args []string) {
		sheets.CreateSheetServiceWithAPIKey(viper.GetString("api-key"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		utils.DisplaySummary(sheets.GetSummary(assignee))
	},
}

func init() {
	sheetCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().StringVar(&assignee, "user", "", "User tracking the enhancements")
}
