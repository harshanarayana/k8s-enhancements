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
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"k8s-enhancements/common"
	"k8s-enhancements/models"
	"k8s-enhancements/utils"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var trackSpec models.TrackSpec
var initialComment bool

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Update tracking status for an Enhancement in Local cache",
	Long:  `Update tracking status for an Enhancement in Local cache`,
	Run: func(cmd *cobra.Command, args []string) {
		basePath := common.GetConfigHome()
		trackFile := strings.Join([]string{basePath, "track.json"}, string(os.PathSeparator))
		tracker := utils.GetTrackingData()
		var spec models.TrackSpec
		issue := viper.GetString("issue")
		status := viper.GetString("status")
		note := viper.GetString("note")

		if item, ok := tracker.Records[issue]; !ok {
			spec.IssueID = issue
			spec.Status = status
			spec.Note = note
			spec.Date = time.Now().Format("January 2, 2006")
		} else {
			spec.IssueID = item.IssueID
			spec.Status = item.Status
			spec.Note = item.Note
			spec.Date = time.Now().Format("January 2, 2006")

			if status != "" {
				spec.Status = status
			}

			if note != "" {
				spec.Note = note
			}
		}

		if viper.GetBool("initial-comment") {
			if spec.Note == "" {
				spec.Note = "First Notification Sent"
			}
			if spec.Status == "" {
				spec.Status = "tracked"
			}
		}

		tracker.Records[issue] = spec
		b, _ := json.Marshal(tracker)
		_ = ioutil.WriteFile(trackFile, b, 0755)
	},
}

func init() {
	sheetCmd.AddCommand(trackCmd)

	trackCmd.PersistentFlags().StringVar(&trackSpec.IssueID, "issue", "", "GitHub Issue ID")
	trackCmd.PersistentFlags().StringVar(&trackSpec.Status, "status", "", "Tracking Status for the Issue")
	trackCmd.PersistentFlags().StringVar(&trackSpec.Note, "note", "", "Tracking Note if any")
	trackCmd.PersistentFlags().BoolVar(&initialComment, "initial-comment", false, "Indicate current tracking is initial comment")

	_ = viper.BindPFlags(trackCmd.PersistentFlags())
}
