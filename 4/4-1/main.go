package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	d := 10 * time.Second
	select {
	case <-time.After(d):
		break
	}

	fmt.Println("end")
}
