package main

import (
	"fmt"
	"time"
)

// author by yijia
// Date is 2019-12-30
// function is about "Chinese proverb: Fish for three days and two days on suning nets", 
//so input any date and caculate this day is on fishing or on holiday, compares to the begging date:19900101

func main() {
	var date string
	fmt.Println("please input the tested date and format like '20060102'")
	fmt.Scanln(&date)
	fmt.Println("Your input date is: ", date)
	totalDay := countDay(date)
	fmt.Println("Until 1990-01-01 total Days are: ", totalDay)
	if totalDay%5 == 4 || totalDay%5 == 0 {
		fmt.Println("input date is on holiday...")
	} else {
		fmt.Println("input date is on fishing...")
	}
}

//Test how many days until 19900101
func countDay(date string) int {
	var totalDay int
	layout := "20060102"
	var perMonth = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	//get time.Time of the input date
	inputDate, _ := time.Parse(layout, date)
	//accumulate totaldays for years
	for year := 1990; year < inputDate.Year(); year++ {
		totalDay += 365 + isLeap(year)
	}
	//Feb IS 29 for leap year
	perMonth[2] += isLeap(inputDate.Year())
	//accumulate month for input date
	for m := 0; m < int(inputDate.Month()); m++ {
		totalDay += perMonth[m]
	}
	//accumulate days for input date
	totalDay += inputDate.Day()
	return totalDay
}

//assert learp year
func isLeap(year int) int {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return 1
	} else {
		return 0
	}
}
