package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Can have a standalone dimension and a value type instead.
type Matrix struct {
	Dimension int
	Values    map[int][]int
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	matrix := &Matrix{}

	dimension, _ := reader.ReadString('\n')

	matrix.Dimension, _ = strconv.Atoi(strings.TrimSpace(dimension))

	matrix.Values = make(map[int][]int, matrix.Dimension)

	for i := 0; i < matrix.Dimension; i++ {
		d, _ := reader.ReadString('\n')

		data := strings.Split(strings.TrimSpace(d), " ")

		//make this else GO panics
		matrix.Values[i] = make([]int, matrix.Dimension)

		for k, v := range data {
			matrix.Values[i][k], _ = strconv.Atoi(v)
		}
	}

	//Add all the data between the "front left" and the "back right" of the first matrix row

	var frontLeftDiagonalSummation, backLeftDiagonalSummation int

	//Move forward
	for i := 0; i < matrix.Dimension; i++ {
		frontLeftDiagonalSummation += matrix.Values[i][i]
	}

	//Move backwards
	for i := 0; i < matrix.Dimension; i++ {
		backLeftDiagonalSummation += matrix.Values[i][matrix.Dimension-1-i]
	}

	sum := frontLeftDiagonalSummation - backLeftDiagonalSummation

	if sum < 0 {
		sum = int(math.Abs(float64(sum)))
	}

	fmt.Println(sum)
}
