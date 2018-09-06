package main

import "fmt"

func main() {
	urlString := BuildURL("GOOGL")
	//fmt.Println(urlString)
	data := Requesting(urlString)
	result, _ := parseRecords(data)

	for _, element := range result {
		fmt.Printf("%v\n\t%f\n\t%f\n\t%f\n\t%f\n\t%f\n\t%f\n\t%d\n\t%f\n\t%f\n", element.Date, element.Open, element.High, element.High, element.Low, element.Close, element.AdjustedClose, element.Volume, element.Dividend, element.Split)
	}
}
