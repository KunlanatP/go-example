package model

type DTOCreateBook struct {
	Name  string  `json:"name"`
	Desc  string  `json:"description"`
	Price float32 `json:"price"`
}

type DTOUpdateBook struct {
	Price float32 `json:"price"`
}

type DTOCreateTheme struct {
	AppID        string `json:"appId"`
	ThemeContent string `json:"themeContent"`
}
