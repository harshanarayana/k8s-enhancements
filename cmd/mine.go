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
	"k8s-enhancements/utils"
)

var userName string
var trackStatus string

// mineCmd represents the mine command
var mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "List all the items in tracking sheet in my name",
	Long:  `Fetch and display all the items in my name from K8s Enhancements Google Sheet`,
	PreRun: func(cmd *cobra.Command, args []string) {
		sheets.CreateSheetServiceWithAPIKey(viper.GetString("api-key"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		utils.DisplayRows(sheets.GetMyAssignmentsV2(viper.GetString("user"), viper.GetString("enhancement-status")))
	},
}

func init() {
	sheetCmd.AddCommand(mineCmd)

	mineCmd.PersistentFlags().StringVar(&userName, "user", "", "Filter values for C10")
	mineCmd.PersistentFlags().StringVar(&trackStatus, "enhancement-status", "", "Enhancement Status to Filter")
	_ = viper.BindPFlags(mineCmd.PersistentFlags())
}
