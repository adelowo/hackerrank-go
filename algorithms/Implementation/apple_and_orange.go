package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var reader = bufio.NewReader(os.Stdin)

	samsHouse, _ := reader.ReadString('\n')

	sam := makeInts(samsHouse)

	t, _ := reader.ReadString('\n')

	trees := makeInts(t)

	n, _ := reader.ReadString('\n')

	numsOfFruitsThatFell := makeInts(n)

	ad, _ := reader.ReadString('\n')

	appleDistance := makeInts(ad)

	od, _ := reader.ReadString('\n')

	orangeDistance := makeInts(od)

	//apples

	type apple map[int]int
	type orange map[int]int

	apples := make(apple, len(appleDistance))
	oranges := make(orange, len(orangeDistance))

	for i := 0; i < numsOfFruitsThatFell[0]; i++ {
		apples[i] = trees[0] + appleDistance[i]
	}

	for i := 0; i < numsOfFruitsThatFell[1]; i++ {
		oranges[i] = trees[1] + orangeDistance[i]
	}

	fmt.Println(countWithinRegion(sam, apples))
	fmt.Println(countWithinRegion(sam, oranges))
}

func makeInts(s string) map[int]int {

	i := make(map[int]int, 8)

	for k, v := range strings.Split(strings.TrimSpace(s), " ") {
		i[k], _ = strconv.Atoi(v)
	}

	return i
}

func countWithinRegion(region, fruits map[int]int) int {

	var count int

	for i := 0; i < len(fruits); i++ {
		if fruits[i] >= region[0] && fruits[i] <= region[1] {
			count++
		}
	}

	return count
}
