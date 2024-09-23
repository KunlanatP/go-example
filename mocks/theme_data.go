package mocks

import "encoding/json"

type ThemeDesign struct {
	ThemeColor ThemeColor `json:"themeColor"`
}

type ThemeColor struct {
	Primary      ThemeMode `json:"primary"`
	OnPrimary    ThemeMode `json:"onPrimary"`
	Secondary    ThemeMode `json:"secondary"`
	OnSecondary  ThemeMode `json:"onSecondary"`
	Error        ThemeMode `json:"error"`
	OnError      ThemeMode `json:"onError"`
	Background   ThemeMode `json:"background"`
	OnBackground ThemeMode `json:"onBackground"`
	Surface      ThemeMode `json:"surface"`
	OnSurface    ThemeMode `json:"onSurface"`
	Text         ThemeMode `json:"text"`
}

type ThemeMode struct {
	LightMode string `json:"lightMode"`
	DarkMode  string `json:"darkMode"`
}

func DefaultDataContent() string {
	themeDesign := ThemeDesign{
		ThemeColor: ThemeColor{
			Primary: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
			OnPrimary: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
			Secondary: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
			OnSecondary: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
			Error: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
			OnError: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
			Background: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
			OnBackground: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
			Surface: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
			OnSurface: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
			Text: ThemeMode{
				LightMode: "#FFF4F4F4",
				DarkMode:  "#FFF4F4F4",
			},
		},
	}

	b, _ := json.Marshal(themeDesign)
	return string(b)
}
