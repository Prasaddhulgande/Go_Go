package main

import (
	"fmt"
	"time"
)

func main() {
	//Time is  apowerfull pkg
	// formats of date & time  yyyy/mm/dd/  hh/mm/ss
	// 02-01-2006 15:04:05 on this date golang officialy launched that's why we need to provide date in this formats only
	// if you wants to date & time owyn ways.

	currentTime := time.Now()

	fmt.Println("Current time is: ", currentTime)
	formattedTime := currentTime.Format("02-01-2006, Monday, 03:04:05 PM")
	fmt.Println("formatted time is: ", formattedTime)

	//If time is in string formates available but we need it in  time formates then...

	layout_str := "02/01/2006"
	datestr := "08/11/1999"

	formate_time, _ := time.Parse(layout_str, datestr)
	fmt.Println("In timeFormatted Time: ", formate_time)

	// add one more day in current time
	new_date := currentTime.Add(2 * time.Hour)
	fmt.Println("new_date Time: ", new_date)
	formated_newDate := new_date.Format("2006/01/02, Monday")
	fmt.Println("formated_newDate Time: ", formated_newDate)

}
