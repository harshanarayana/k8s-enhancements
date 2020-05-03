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
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Helper Utilities to manage kubernetes/enhancements on GitHub",
	Long:  `Helper Utilities to manage kubernetes/enhancements on GitHub`,
}

func init() {
	rootCmd.AddCommand(gitCmd)

	gitCmd.PersistentFlags().StringVarP(&gitAccess.Owner, "owner", "o", "kubernetes", "GitHub Repo Owner")
	gitCmd.PersistentFlags().StringVarP(&gitAccess.Repo, "repo", "r", "enhancements", "GitHub Repository")
	gitCmd.PersistentFlags().StringVarP(&gitAccess.UserName, "git-user", "u", "", "GitHub Username")
	gitCmd.PersistentFlags().StringVarP(&gitAccess.AccessToken, "git-access-token", "t", "", "GitHub Access token")

	_ = gitCmd.MarkPersistentFlagRequired("owner")
	_ = gitCmd.MarkPersistentFlagRequired("repo")

	_ = viper.BindPFlags(gitCmd.PersistentFlags())
}
