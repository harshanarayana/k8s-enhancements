package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"k8s-enhancements/common"
	"os"
	"strings"
)

func GetCommentMessage(templateName string, mentions, missing, stages []string) string {
	basePath := common.GetConfigHome()
	templatePath := strings.Join([]string{basePath, "templates", templateName}, string(os.PathSeparator))

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		panic(err)
	} else {
		tmpl := template.Must(template.ParseFiles(templatePath))
		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, struct {
			Mentions string
			Missing  string
			Stage    string
		}{
			Mentions: strings.Join(getUsers(mentions), " / "),
			Missing:  getMissing(missing),
			Stage:    getStage(stages),
		}); err != nil {
			panic(err)
		}
		return tpl.String()
	}
}

func getStage(stages []string) string {
	var fStage []string
	for _, s := range stages {
		fStage = append(fStage, fmt.Sprintf("`%s`", s))
	}
	return strings.Join(fStage, "/")
}

func getUsers(mentions []string) []string {
	var users []string = make([]string, 0)
	for _, m := range mentions {
		users = append(users, "@"+m)
	}
	return users
}

func getMissing(missing []string) string {
	missing = mapStrings(missing, func(v string) string {
		return fmt.Sprintf("`%s`", v)
	})

	if len(missing) > 2 {
		lastString := missing[len(missing)-1]
		return strings.Join([]string{strings.Join(missing[:len(missing)-1], ", "), lastString}, " and ")
	} else if len(missing) > 1 {
		return strings.Join(missing, " and ")
	} else {
		return strings.Join(missing, "")
	}
}

func mapStrings(stringArray []string, f func(v string) string) []string {
	mappedStrings := make([]string, 0)
	for _, s := range stringArray {
		mappedStrings = append(mappedStrings, f(s))
	}
	return mappedStrings
}
