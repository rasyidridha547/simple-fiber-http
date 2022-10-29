package model

type Helicopter []struct {
	Name string `json:"name"`
	Size string `json:"size"`
	Year int32  `json:"year"`
	Type string `json:"type"`
}
