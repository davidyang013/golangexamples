package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Unix()
	fmt.Println(t)

	fmt.Println(time.Unix(t, 0).String())

	t = time.Now().UnixNano()
	fmt.Println(t) //nano time

	fmt.Println(time.Now().String())

	fmt.Println(time.Now().Format("2006year 01month 02day"))

	p := time.Now()
	fmt.Println(p.Weekday().String())
}
