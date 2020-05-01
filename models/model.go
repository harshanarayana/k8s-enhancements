package models

type GitAccess struct {
	UserName    string
	AccessToken string
	Owner       string
	Repo        string
}

type EnhancementRow struct {
	IssueID           string
	EnhancementTitle  string
	EnhancementStatus string
	StageStatus       string
	Stage             string
	Sig               string
	KEPOwner          string
	ProposalLink      string
	KEPState          string
	KKPRs             string
	PRList            string
	Notes             string
}

var AttributeMapper = map[int]string{
	0:  "IssueID",
	1:  "EnhancementTitle",
	3:  "EnhancementStatus",
	4:  "StageStatus",
	5:  "Stage",
	6:  "Sig",
	7:  "KEPOwner",
	8:  "ProposalLink",
	9:  "KEPState",
	10: "KKPRs",
	11: "PRList",
	12: "Notes",
}

type TrackSpec struct {
	IssueID string `json:"issue_id"`
	Status  string `json:"status"`
	Note    string `json:"note"`
	Date    string `json:"date"`
}

type Tracker struct {
	Records map[string]TrackSpec `json:"records"`
}

type SpreadSheetCell struct {
	Row    int
	Column string
}

var TypeToColumnMap = map[string]string{
	"status":      "D",
	"stageStatus": "E",
	"stage":       "F",
	"notes":       "M",
}

type Summary struct{
	Count int
	IssueData map[string]string
}