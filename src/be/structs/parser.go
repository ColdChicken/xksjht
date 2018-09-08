package structs

type ParserContentBlock struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Link    string `json:"link"`
}

type ParserContent struct {
	Type      string                `json:"type"`
	SectionID string                `json:"section_id"`
	Content   string                `json:"content"`
	Source    string                `json:"source"`
	Value     string                `json:"value"`
	Ordered   bool                  `json:"ordered"`
	Contents  []*ParserContentBlock `json:"contents"`
}

type ParserResult struct {
	Title    string           `json:"title"`
	Contents []*ParserContent `json:"contents"`
}
