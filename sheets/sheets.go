package sheets

import (
	"fmt"
	"google.golang.org/api/sheets/v4"
	"k8s-enhancements/models"
	"reflect"
	"strconv"
	"strings"
)

const (
	K8SSheetId      = "1E89GNZRwmnfVerOqWqrmzVII8TQlQ0iiI_FgYKxanGs"
	EnhancementsTab = "Enhancements!A11:M"
)

func GetMyAssignmentsV2(username, status string) []*models.EnhancementRow {
	enhancements := make([]*models.EnhancementRow, 0)
	if resp, err := GetSheetService().Spreadsheets.Values.Get(K8SSheetId, EnhancementsTab).Do(); err != nil {
		panic(err)
	} else {
		if len(resp.Values) < 1 {
			return enhancements
		} else {
			for _, row := range resp.Values {
				if username != "" && fmt.Sprintf("%s", row[2]) != username {
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

func findCell(issueID, infoType string) models.SpreadSheetCell {
	var cell models.SpreadSheetCell

	if infoType != "" {
		if col, ok := models.TypeToColumnMap[infoType]; !ok {
			panic(fmt.Errorf("failed to determine column for entity type %s", infoType))
		} else {
			cell.Column = col
		}
	}

	if resp, err := GetSheetService().Spreadsheets.Values.Get(K8SSheetId, EnhancementsTab).Do(); err != nil {
		panic(err)
	} else {
		if len(resp.Values) < 1 {
			return cell
		} else {
			currentRow := 10
			for _, row := range resp.Values {
				currentRow += 1
				if fmt.Sprintf("%s", row[0]) == issueID {
					cell.Row = currentRow
					break
				}
			}
		}
	}
	return cell
}

func UpdateRecord(issueId string, data map[string]interface{}) {
	cell := findCell(issueId, "")
	for infoType, value := range data {
		if col, ok := models.TypeToColumnMap[infoType]; !ok {
			panic(fmt.Errorf("failed to determine column for type %s", infoType))
		} else {
			cell.Column = col
		}
		cellRange := strings.Join([]string{"Enhancements!", cell.Column, strconv.Itoa(cell.Row)}, "")
		if _, err := GetSheetService().Spreadsheets.Values.Update(K8SSheetId, cellRange, &sheets.ValueRange{
			MajorDimension: "COLUMNS",
			Values:         [][]interface{}{{value}},
		}).ValueInputOption("RAW").Do(); err != nil {
			panic(err)
		}
	}
}

func GetSummary(owner string) map[string]models.Summary {
	items := GetMyAssignmentsV2(owner, "")
	data := make(map[string]models.Summary)

	for _, item := range items {

		if _, ok := data[item.EnhancementStatus]; !ok {
			data[item.EnhancementStatus] = models.Summary{IssueData: map[string][]models.IssueSummary{}}
		}

		if _, ok := data[item.EnhancementStatus].IssueData[item.Stage]; !ok {
			data[item.EnhancementStatus].IssueData[item.Stage] = make([]models.IssueSummary, 0)
		}

		data[item.EnhancementStatus].IssueData[item.Stage] = append(data[item.EnhancementStatus].IssueData[item.Stage], models.IssueSummary{
			Title:   item.EnhancementTitle,
			Stage:   item.Stage,
			Status:  item.StageStatus,
			IssueId: item.IssueID,
		})
	}
	return data
}
