package payloads

type RepoRequest struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

type ScanRequest struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type GenericResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message *string     `json:"message,omitempty"`
}

type ListRepoResponse struct {
	Page        int                     `json:"page"`
	ItemPerPage int                     `json:"item_count"`
	TotalCount  int                     `json:"total_count"`
	ItemList    []*ListRepoResponseItem `json:"item_list"`
}

type ListRepoResponseItem struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	URL        string  `json:"url"`
	ScanStatus *string `json:"scan_status"`
	Timestamp  *string `json:"timestamp"`
}

type ViewRepoResponse struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	URL        string     `json:"url"`
	ScanStatus *string    `json:"scan_status"`
	Timestamp  *string    `json:"timestamp"`
	Findings   []*Finding `json:"findings"`
}

type Finding struct {
	Type     string          `json:"type"`
	RuleID   string          `json:"ruleId"`
	Location FindingLocation `json:"location"`
	Metadata FindingMetadata `json:"metadata"`
}

type FindingLocation struct {
	Path      string                  `json:"path"`
	Positions FindingLocationPosition `json:"positions"`
}

type FindingLocationPosition struct {
	Begin FindingLocationPositionBegin `json:"begin"`
}

type FindingLocationPositionBegin struct {
	Line int `json:"line"`
}

type FindingMetadata struct {
	Description string `json:"description"`
	Severity    string `json:"severity"`
}
