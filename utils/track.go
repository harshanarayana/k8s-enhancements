package utils

import (
	"encoding/json"
	"io/ioutil"
	"k8s-enhancements/common"
	"k8s-enhancements/models"
	"os"
	"strings"
)

func GetTrackingData() models.Tracker {
	var tracker models.Tracker
	basePath := common.GetConfigHome()
	trackFile := strings.Join([]string{basePath, "track.json"}, string(os.PathSeparator))
	if _, err := os.Stat(trackFile); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		tracker.Records = make(map[string]models.TrackSpec, 0)
	} else {
		if b, err := ioutil.ReadFile(trackFile); err != nil {
			panic(err)
		} else {
			if err := json.Unmarshal(b, &tracker); err != nil {
				panic(err)
			}
		}
	}
	return tracker
}
