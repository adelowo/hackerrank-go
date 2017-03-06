package main

import (
	"fmt"
	"strings"
)

func main() {
	var numberOfStairs int

	fmt.Scanf("%d", &numberOfStairs)

	for i := 1; i <= numberOfStairs; i++ {
		fmt.Printf("%s%s\n", strings.Repeat(" ", numberOfStairs-i), strings.Repeat("#", i))
	}
}
