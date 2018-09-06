package main

/*
	A struct is only necessary right now, I will expand
	it to more API functions if the need arises
*/

import (
	"time"
)

// TimeSeriesAD stores all the Data received from the TIME_SERIES_DAILY_ADJUSTED API function
type TimeSeriesAD struct {
	Date          time.Time `json:"date"`
	Open          float64   `json:"open"`
	High          float64   `json:"high"`
	Low           float64   `json:"low"`
	Close         float64   `json:"close"`
	AdjustedClose float64   `json:"adjustedClose"`
	Volume        int       `json:"volume"`
	Dividend      float64   `json:"dividend"`
	Split         float64   `json:"split"`
}
