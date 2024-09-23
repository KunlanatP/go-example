package domain

type Books struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Desc  string  `json:"description"`
	Price float32 `json:"price"`
}
