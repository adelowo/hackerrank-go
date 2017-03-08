package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const WIDTH = 1

func main() {

	var reader = bufio.NewReader(os.Stdin)

	h, _ := reader.ReadString('\n')

	wordHeights := make(map[int]int, 26)

	for k, v := range strings.Split(strings.TrimSpace(h), " ") {
		i, _ := strconv.Atoi(v)
		wordHeights[k] = i
	}

	w, _ := reader.ReadString('\n')

	var highest int

	for _, v := range strings.TrimSpace(w) {
		temp := wordHeights[int(v-rune('a'))]

		if temp > highest {
			highest = temp
		}
	}

	fmt.Println((len(strings.Split(strings.TrimSpace(w), "")) * WIDTH) * highest)
}
