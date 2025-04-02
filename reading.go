package main

type ReadingRequest struct {
	Value float64 `json:"value" form:"value" xml:"value"`
	Epoch int64   `json:"epoch" form:"epoch" xml:"epoch"`
}
