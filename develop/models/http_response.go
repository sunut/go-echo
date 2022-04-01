package models

type StatusCode struct {
	Code        int    `json:"code" xml:"code"`
	Description string `json:"description" xml:"description"`
}

type Status struct {
	Status StatusCode `json:"status" xml:"status"`
}
