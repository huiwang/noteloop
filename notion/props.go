package notion

type TitleProp struct {
	Type  string  `json:"type"`
	Title []Title `json:"title"`
}

type Title struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

// create title prop with a single title
func CreateTitleProp(title string) TitleProp {
	return TitleProp{
		Type: "title",
		Title: []Title{
			{
				Type: "text",
				Text: Text{
					Content: title,
				},
			},
		},
	}
}

type NumberProp struct {
	Type   string  `json:"type"`
	Number float64 `json:"number"`
}

// create number prop with a single number
func CreateNumberProp(number float64) NumberProp {
	return NumberProp{
		Type:   "number",
		Number: number,
	}
}

type DateProp struct {
	Type string `json:"type"`
	Date Date   `json:"date"`
}
type Date struct {
	Start string `json:"start"`
}

// create date prop with a single date
func CreateDateProp(date string) DateProp {
	return DateProp{
		Type: "date",
		Date: Date{
			Start: date,
		},
	}
}
