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
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"k8s-enhancements/common"
	"os"
	"strings"
)

var tmplName string

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new GitHub issue comment template",
	Long:  `Create new GitHub issue comment template`,
	Run: func(cmd *cobra.Command, args []string) {
		tmplName = viper.GetString("template")
		if tmplName == "" {
			panic(fmt.Errorf("please enter a valid --template arg"))
		}

		templateFile := strings.Join([]string{common.GetConfigHome(), "templates", tmplName}, string(os.PathSeparator))
		if _, err := os.Stat(templateFile); err == nil || !os.IsNotExist(err) {
			panic(fmt.Errorf("template file %s already exists", templateFile))
		}
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter Template Body [Terminate the message with '.done']: ")
		var body = make([]string, 0)
		for scanner.Scan() {
			d := scanner.Text()
			if d == ".done" {
				break
			}
			body = append(body, d)
		}

		_ = ioutil.WriteFile(templateFile, []byte(strings.Join(body, "\n")), 0755)
	},
}

func init() {
	templatesCmd.AddCommand(newCmd)
}
