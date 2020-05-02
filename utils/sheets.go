package utils

import (
	"github.com/olekukonko/tablewriter"
	"k8s-enhancements/models"
	"os"
	"sort"
	"strconv"
	"strings"
)

var tracker models.Tracker

func DisplaySummary(summary map[string]models.Summary) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Enhancement Status", "Issue Count", "Issue ID", "Issue Description"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)

	var data [][]string

	for status, info := range summary {
		issueIDs := make([]string, 0)
		for id, _ := range info.IssueData {
			issueIDs = append(issueIDs, id)
		}

		// Pass in our list and a func to compare values
		sort.Slice(issueIDs, func(i, j int) bool {
			numA, _ := strconv.Atoi(issueIDs[i])
			numB, _ := strconv.Atoi(issueIDs[j])
			return numA < numB
		})

		for _, id := range issueIDs {
			r := make([]string, 0)
			r = append(r, status)
			r = append(r, strconv.Itoa(info.Count))
			r = append(r, id)
			r = append(r, info.IssueData[id])
			data = append(data, r)
		}
	}

	table.SetAutoMergeCells(true)
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColWidth(80)
	table.SetAutoWrapText(true)
	table.AppendBulk(data)
	table.Render()
}

func DisplayRows(issues []*models.EnhancementRow) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Issue ID", "Issue Title", "Enhancement Status", "Stage Status", "Stage", "Sig", "KEP Owner", "Note", "Updated Date", "Tracking"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)

	var data [][]string

	for _, issue := range issues {
		date, note, state := getTrackingState(issue.IssueID)
		data = append(
			data,
			[]string{
				issue.IssueID,
				issue.EnhancementTitle,
				issue.EnhancementStatus,
				issue.StageStatus,
				issue.Stage,
				issue.Sig,
				issue.KEPOwner,
				note,
				date,
				state,
			},
		)
	}
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColWidth(52)
	table.SetColMinWidth(1, 50)
	table.SetColMinWidth(9, 10)
	table.SetAutoWrapText(true)
	table.AppendBulk(data)
	table.Render()
}

func transformTrackingStatus(status string) string {
	switch strings.ToLower(status) {
	case "tracked", "track", "t":
		return "︎✔︎"
	case "ignored", "ignore", "i":
		return "✘"
	case "skip", "s":
		return "►"
	case "info":
		return "ℹ"
	case "watch", "monitor", "w", "m":
		return "◎"
	default:
		return "𖡄"
	}
}

func getTrackingState(issue string) (string, string, string) {
	if item, ok := tracker.Records[issue]; !ok {
		return "", "", transformTrackingStatus("unknown")
	} else {
		return item.Date, item.Note, transformTrackingStatus(item.Status)
	}
}

func init() {
	tracker = GetTrackingData()
}
