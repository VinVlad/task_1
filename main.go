package main

import (
	"fmt"
	"time"
)

func main() {

	//layout := "02.01.2006 15:04"

	now := time.Now()
	//now := "2024-01-06 21:11:37.2872925 +1000 +10 m=+0.001561901"
	//time.Parse(layout, now)

	//Day := now.Day()
	//Month := now.Month()
	//Year := now.Year()

	fmt.Println("Hello!")
	//fmt.Printf("%v.%v.%v", Day, Month, Year)

	fmt.Println(now)
}
