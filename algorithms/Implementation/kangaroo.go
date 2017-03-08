package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type kangaroo [2]int

func main() {

	kanaroos, _ := bufio.NewReader(os.Stdin).
		ReadString('\n')

	data := strings.Split(strings.TrimSpace(kanaroos), " ")

	k1 := makeKangaroo(data[0], data[1])
	k2 := makeKangaroo(data[2], data[3])

	if k1[1] > k2[1] && ((k2[0]-k1[0])%(k1[1]-k2[1]) == 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func makeKangaroo(start string, jumpRate string) kangaroo {

	var k kangaroo

	k[0], _ = strconv.Atoi(start)
	k[1], _ = strconv.Atoi(jumpRate)

	return k

}
