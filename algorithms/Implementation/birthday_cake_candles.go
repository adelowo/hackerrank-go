package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func birthdayCakeCandles(ar []int) int {

	var count = 0

	sort.Ints(ar)

	var max = ar[len(ar)-1]

	for _, v := range ar {
		if v == max {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	arCount, err := strconv.ParseInt(strings.TrimSpace(age), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	candlesHeight, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	arTemp := strings.Split(strings.TrimSpace(candlesHeight), " ")

	var ar []int

	for i := 0; i < int(arCount); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int(arItemTemp)
		ar = append(ar, arItem)
	}

	result := birthdayCakeCandles(ar)

	fmt.Fprintf(os.Stdout, "%d\n", result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
