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
	"github.com/harshanarayana/k8s-enhancements/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ListOptions struct {
	MaxSize  int
	Milestone string
	State string
	Assignee string
	Labels []string
}

var options ListOptions

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available Issues on Enhancement Repo",
	Long: `List Enhancements under kubernetes/enhancements repo based on search criteria specified`,
	Run: func(cmd *cobra.Command, args []string) {
		git.ListIssues(viper.GetString("milestone"), viper.GetString("state"), viper.GetString("assignee"), viper.GetStringSlice("labels"), viper.GetInt("max-size"))
	},
}

func init() {
	issuesCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().IntVarP(&options.MaxSize, "max-size", "m", 0, "Max Records to display with paginated request")
	listCmd.PersistentFlags().StringVar(&options.Milestone, "milestone", "", "Milestone assigned to the Issue")
	listCmd.PersistentFlags().StringVar(&options.State, "state", "", "Issue State")
	listCmd.PersistentFlags().StringVar(&options.Assignee, "assignee", "", "Assigned User")
	listCmd.PersistentFlags().StringArrayVar(&options.Labels, "label", []string{}, "Labels Assigned to the Issue")
	_ = viper.BindPFlags(listCmd.PersistentFlags())
}
