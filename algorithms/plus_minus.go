package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type array []int

func main() {

	var reader = bufio.NewReader(os.Stdin)

	size, _ := reader.ReadString('\n')

	arraySize, _ := strconv.Atoi(strings.TrimSpace(size))

	a, _ := reader.ReadString('\n')

	ar := strings.Split(strings.TrimSpace(a), " ")

	var arr array

	for i := 0; i < arraySize; i++ {
		d, _ := strconv.Atoi(ar[i])
		arr = append(arr, d)
	}

	var positives, negatives, neutrals int

	for i := 0; i < arraySize; i++ {
		if arr[i] == 0 {
			neutrals++
		}

		if arr[i] > 0 {
			positives++
		}

		if arr[i] < 0 {
			negatives++
		}
	}

	arrLength := len(arr)

	fmt.Printf("%.6f \n", float64(positives)/float64(arrLength))
	fmt.Printf("%.6f \n", float64(negatives)/float64(arrLength))
	fmt.Printf("%.6f \n", float64(neutrals)/float64(arrLength))
}
