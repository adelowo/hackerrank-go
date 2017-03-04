package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var reader = bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	da := strings.Split(strings.TrimSpace(input), " ")

	var data []int

	for _, v := range da {
		digit, _ := strconv.Atoi(v)
		data = append(data, digit)
	}

	var result []int

	for i := 0; i < len(data); i++ {
		var temp int

		for y := 0; y < len(data); y++ {
			if y == i {
				continue
			}

			temp += data[y]
		}

		result = append(result, temp)
	}

	sort.Ints(result)

	fmt.Printf("%d %d\n", result[0], result[len(result)-1])
}
