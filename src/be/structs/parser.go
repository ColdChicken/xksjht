package structs

type ParserContentBlock struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Link    string `json:"link"`
	RefIdx  int64  `json:"refIdx"`
}

type ParserContent struct {
	Type         string                `json:"type"`
	SectionID    string                `json:"section_id"`
	SectionLevel int64                 `json:"section_level"`
	Content      string                `json:"content"`
	Source       string                `json:"source"`
	Value        string                `json:"value"`
	Ordered      bool                  `json:"ordered"`
	Contents     []*ParserContentBlock `json:"contents"`
}

type ParserResult struct {
	Title    string           `json:"title"`
	Contents []*ParserContent `json:"contents"`
}
