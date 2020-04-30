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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s-enhancements/sheets"
)

var updates map[string]string
var gitIssue string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Google Spreadsheet with Details",
	Long: `Update Google Spreadsheet with Details`,
	PreRun: func(cmd *cobra.Command, args []string) {
		sheets.CreateSheetServiceWithOAuth()
	},
	Run: func(cmd *cobra.Command, args []string) {
		data := make(map[string]interface{}, 0)

		for t, v := range updates {
			data[t] = v
		}
		sheets.UpdateRecord(gitIssue, data)
	},
}

func init() {
	sheetCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringToStringVar(&updates, "records", nil, "Specify fields to update with respective values")
	updateCmd.Flags().StringVar(&gitIssue, "eid", "", "Enhancement ID/GitHub Issue ID")
	_ = viper.BindPFlags(updateCmd.PersistentFlags())
}
