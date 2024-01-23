package notion

type Parent struct {
	Type       string `json:"type"`
	DatabaseID string `json:"database_id"`
}

type Grocery struct {
	Parent     Parent       `json:"parent"`
	Properties GroceryProps `json:"properties"`
}

type GroceryProps struct {
	GroceryItem TitleProp  `json:"Grocery item"`
	Price       NumberProp `json:"Price"`
	LastOrdered DateProp   `json:"Last ordered"`
}

func Create(name string, price float64, date string) *Grocery {
	return &Grocery{
		Parent: Parent{
			Type:       "database_id",
			DatabaseID: "0b4eb47acb29438597538cc4b4d05ef3",
		},
		Properties: GroceryProps{
			GroceryItem: CreateTitleProp(name),
			Price:       CreateNumberProp(price),
			LastOrdered: CreateDateProp(date),
		},
	}
}
