package common

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"strings"
)

func GetConfigHome() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Join([]string{home, ".k8s-enhancements"}, string(os.PathSeparator))
}
