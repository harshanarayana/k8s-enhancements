package utils

import (
	"bytes"
	"html/template"
	"k8s-enhancements/common"
	"os"
	"strings"
)

func GetCommentMessage(templateName string, mentions []string) string {
	basePath := common.GetConfigHome()
	templatePath := strings.Join([]string{basePath, "templates", templateName}, string(os.PathSeparator))

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		panic(err)
	} else {
		tmpl := template.Must(template.ParseFiles(templatePath))
		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, struct {
			Mentions string
		}{
			Mentions: strings.Join(getUsers(mentions), " / "),
		}); err != nil {
			panic(err)
		}
		return tpl.String()
	}
}

func getUsers(mentions []string) []string {
	var users []string = make([]string, 0)
	for _, m := range mentions {
		users = append(users, "@" + m)
	}
	return users
}