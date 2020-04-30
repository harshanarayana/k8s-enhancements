package sheets

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"k8s-enhancements/models"
	"reflect"
)

var sheetService *sheets.Service

const (
	K8SSheetId      = "1E89GNZRwmnfVerOqWqrmzVII8TQlQ0iiI_FgYKxanGs"
	EnhancementsTab = "Enhancements!A11:M"
)

func Setup(apiKey string)  {
	if svc, err := sheets.NewService(context.Background(), option.WithAPIKey(apiKey)); err != nil {
		panic(err)
	} else {
		sheetService = svc
	}
}

func GetMyAssignmentsV2(username, status string) []*models.EnhancementRow {
	enhancements := make([]*models.EnhancementRow, 0)
	if resp, err := sheetService.Spreadsheets.Values.Get(K8SSheetId, EnhancementsTab).Do(); err != nil {
		panic(err)
	} else {
		if len(resp.Values) < 1 {
			return enhancements
		} else {
			for _, row := range resp.Values {
				if fmt.Sprintf("%s", row[2]) != username {
					continue
				}

				if status != "" && fmt.Sprintf("%s", row[3]) != status {
					continue
				}

				enh := &models.EnhancementRow{}
				if len(row) < 5 {
					continue
				}
				for _, index := range []int{0, 1, 3, 4, 5, 6, 7, 8, 9} {
					reflect.ValueOf(enh).Elem().FieldByName(models.AttributeMapper[index]).SetString(fmt.Sprintf("%s", row[index]))
				}
				enhancements = append(enhancements, enh)
			}
		}
	}
	return enhancements
}
