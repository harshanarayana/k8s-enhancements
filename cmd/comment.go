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
	"k8s-enhancements/git"
	"k8s-enhancements/utils"

	"github.com/spf13/cobra"
)

var issueId int
var template string
var refer []string

// commentCmd represents the comment command
var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Add a Comment to GitHub Issue using Template",
	Long:  `Add a Comment to GitHub Issue using Template`,
	PreRun: func(cmd *cobra.Command, args []string) {
		git.InitGit(viper.GetString("git-access-token"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		message := utils.GetCommentMessage(viper.GetString("tpl"), viper.GetStringSlice("mention"))
		git.AddComment(viper.GetString("owner"), viper.GetString("repo"), message, viper.GetInt("git-issue"))
	},
}

func init() {
	issuesCmd.AddCommand(commentCmd)

	commentCmd.PersistentFlags().IntVar(&issueId, "git-issue", 1, "GitHub Issue ID")
	commentCmd.PersistentFlags().StringVar(&template, "tpl", "initial", "Comment Message template to use")
	commentCmd.PersistentFlags().StringSliceVar(&refer, "mention", []string{}, "Mention use to be listed in the Comment")
	_ = viper.BindPFlags(commentCmd.PersistentFlags())
}
