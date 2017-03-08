package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type studentGrades []int

func main() {

	reader := bufio.NewReader(os.Stdin)

	size, _ := reader.ReadString('\n')

	studentsTotal, _ := strconv.Atoi(strings.TrimSpace(size))

	var grades studentGrades

	for i := 0; i < studentsTotal; i++ {

		g, _ := reader.ReadString('\n')

		grade, _ := strconv.Atoi(strings.TrimSpace(g))

		grades = append(grades, grade)
	}

	for _, v := range gradeByProfessorSamRules(grades) {
		fmt.Println(v)
	}

}

func gradeByProfessorSamRules(grade studentGrades) studentGrades {

	var newGrades studentGrades

	for i := 0; i < len(grade); i++ {
		score := grade[i]

		remainder := score % 5

		if score < 38 || remainder < 3 {
			newGrades = append(newGrades, score)
			continue
		}

		newGrades = append(newGrades, score+5-remainder)
	}

	return newGrades
}
