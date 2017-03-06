package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	arrSize, _ := reader.ReadString('\n')

	data, _ := reader.ReadString('\n')

	//make an array with the size of arrSize

	size, _ := strconv.Atoi(strings.TrimSpace(arrSize))

	arr := make([]int, size)

	d := strings.Split(strings.TrimSpace(data), " ")

	//Fill in the arr with data
	//Force convert them into ints

	answer := 0

	for index, val := range d {
		arr[index], _ = strconv.Atoi(val)
	}

	for _, i := range arr {
		answer += i
	}

	fmt.Println(answer)
}
