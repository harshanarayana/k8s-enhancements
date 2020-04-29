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
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var credentialFile string
// sheetCmd represents the sheet command
var sheetCmd = &cobra.Command{
	Use:   "sheet",
	Short: "Helper Utilities to manage kubernetes/enhancements on Google Sheets",
	Long: `Helper Utilities to manage kubernetes/enhancements on Google Sheets`,
}

func init() {
	rootCmd.AddCommand(sheetCmd)

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	homeConfigPath := strings.Join([]string{home, ".k8s-enhancements"}, string(os.PathSeparator))

	configFile := strings.Join([]string{homeConfigPath, "sheet-credentials.json"}, string(os.PathSeparator))
	sheetCmd.PersistentFlags().StringVar(&credentialFile, "sheet-credentials", configFile, "Google Sheets Access Credentials")

	_ = viper.BindPFlags(sheetCmd.PersistentFlags())
}
