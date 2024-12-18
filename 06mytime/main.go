package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang!")
	presentTime := time.Now()
	fmt.Println("Time: ", presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))
	createdDate := time.Date(2025, time.April, 29, 15, 15, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
