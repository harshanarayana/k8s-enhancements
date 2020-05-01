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
	"k8s-enhancements/common"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var templateName string

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a GitHub Issue Comment Template",
	Long:  `Remove a GitHub Issue Comment Template`,
	Run: func(cmd *cobra.Command, args []string) {
		if templateName != "" {
			basePath := strings.Join([]string{common.GetConfigHome(), "templates", templateName}, string(os.PathSeparator))
			_ = os.Remove(basePath)
		}
	},
}

func init() {
	templatesCmd.AddCommand(rmCmd)

	rmCmd.PersistentFlags().StringVar(&templateName, "template", "", "Template name to cleanup")

	_ = viper.BindPFlags(rmCmd.PersistentFlags())
}
