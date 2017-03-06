package main

import (
	"fmt"
	"time"
)

func main() {

	var input string

	fmt.Scanf("%s", &input)

	//formatted,_ := time.Parse(time.Kitchen, input)
	formatted, _ := time.Parse("3:04:05PM", input)

	fmt.Println(formatted.Format("15:04:05"))
}
