package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var size, data string

	reader := bufio.NewReader(os.Stdin)

	size, _ = reader.ReadString('\n')
	data, _ = reader.ReadString('\n')

	arrSize, _ := strconv.Atoi(size)

	d := make(map[int]int, arrSize)

	for k, v := range strings.Split(strings.TrimSpace(data), " ") {
		d[k], _ = strconv.Atoi(v)
	}

	sum := 0

	for i := 0; i < len(d); i++ {
		sum += d[i]
	}

	fmt.Println(sum)

}
