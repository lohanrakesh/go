package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")

	persentTime := time.Now()

	fmt.Println(persentTime)

	fmt.Println(persentTime.Format("01-02-2006"))

	createdTime := time.Date(1990, time.June, 04, 12, 55, 0, 0, time.UTC)

	fmt.Println(createdTime)

	fmt.Println(createdTime.Format("01-02-2006 Monday"))

}
