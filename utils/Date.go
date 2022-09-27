package utils

import (
	"fmt"
	"time"
)

func CountDays(start_date, end_date string) (int, error) {

	//format := "2006-01-02 15:04:05"

	sDate, error := time.Parse(time.RFC3339, start_date)
	if error != nil {
		fmt.Println("unable to parse Start Date: ", error)
		return 0, error
	}

	eDate, error := time.Parse(time.RFC3339, end_date)
	if error != nil {
		fmt.Println("unable to parse End Date: ", error)
		return 0, error
	}

	diff := eDate.Sub(sDate)

	return int(diff.Hours() / 24), nil
}
