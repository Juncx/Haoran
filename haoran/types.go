package haoran

type ParagraphDetial struct {
	CId     int    `json:"chapterId"`
	PId     int    `json:"paragraphId"`
	Chapter string `json:"chapter"`
	Content string `json:"content"`
	Summary string `json:"summary"`
}

type Chapter struct {
	CId        int               `json:"chapterId"`
	Chapter    string            `json:"chapter"`
	Paragraphs []ParagraphDetial `json:"paragraphs"`
}

type LunYu struct {
	Chapter    string   `json:"chapter"`
	Paragraphs []string `json:"paragraphs"`
}
