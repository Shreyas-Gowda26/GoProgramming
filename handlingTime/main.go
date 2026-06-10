package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Handling time in golang!")

	presentTime := time.Now()
	fmt.Println("Present Time:", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05")) //This is the standard format always follow this!
	createDate := time.Date(2026, time.May, 11, 9, 0, 0, 0, time.UTC)
	fmt.Println("Create Date:", createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday 15:04:05")) //This is the standard format always follow this!
}
