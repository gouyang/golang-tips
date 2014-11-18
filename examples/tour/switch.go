package main

import (
	"fmt"
	"runtime"
	"time"
)

func GuessOS() {
	fmt.Print("GO runs on ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Liunx.")
	case "darwin":
		fmt.Println("OS X.")
	default:
		fmt.Println("%s", os)
	}
}

func GuessDay() {
	fmt.Println("When is Friday?")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("The day after tomorrow.")
	case today + 3:
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}
}

func main() {
	GuessOS()
	GuessDay()
}
