package main

import "fmt"

func main() {
	var week = map[int]string{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 7: "Sunday"}
	for index, value := range week {
		fmt.Println("Day", index, "is", value)
	}
	// _, isPresentSecond = map1[2]
	// _, isPresent
}
