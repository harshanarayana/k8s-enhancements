package utils

import (
	"github.com/olekukonko/tablewriter"
	"k8s-enhancements/models"
	"os"
)

func DisplayRows(issues []*models.EnhancementRow) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Issue ID", "Issue Title", "Enhancement Status", "Stage Status", "Stage", "Sig", "KEP Owner"})

	table.SetHeaderColor(
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
			},
		)
	}
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColMinWidth(1, 75)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()

}