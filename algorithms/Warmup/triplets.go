package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ratings [3]int

type comparisonScore int

func main() {

	var aliceRatings, bobRatings ratings

	var aliceComparisonScores, bobComparisonScores comparisonScore

	var reader = bufio.NewReader(os.Stdin)

	alice, _ := reader.ReadString('\n')

	bob, _ := reader.ReadString('\n')

	aliceRatings = makeRatings(strings.TrimSpace(alice))

	bobRatings = makeRatings(strings.TrimSpace(bob))

	for i := 0; i < len(aliceRatings); i++ {

		if aliceRatings[i] == bobRatings[i] {
			continue
		}

		if aliceRatings[i] > bobRatings[i] {
			aliceComparisonScores += 1
		}

		if aliceRatings[i] < bobRatings[i] {
			bobComparisonScores += 1
		}
	}

	fmt.Println(aliceComparisonScores, bobComparisonScores)

}

func makeRatings(data string) ratings {

	var r ratings

	for k, v := range strings.Split(data, " ") {
		r[k], _ = strconv.Atoi(v)
	}

	return r
}
