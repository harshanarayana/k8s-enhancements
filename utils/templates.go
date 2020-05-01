package utils

import (
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"k8s-enhancements/common"
	"os"
	"path/filepath"
	"strings"
)

func ListTemplates() {
	templatePath := strings.Join([]string{common.GetConfigHome(), "templates"}, string(os.PathSeparator))

	var tmpInfo = make(map[string]string, 0)

	if err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {

			if b, err := ioutil.ReadFile(path); err != nil {
				return nil
			} else {
				tmpInfo[info.Name()] = string(b)
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Template Name", "Template"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)

	var data [][]string

	for template, content := range tmpInfo {
		data = append(data, []string{
			template,
			content,
		})
	}

	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()
}
