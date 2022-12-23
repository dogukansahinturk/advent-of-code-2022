package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var question = `It seems like there is still quite a bit of duplicate work planned. Instead, the Elves would like to know the number of pairs that overlap at all.

In the above example, the first two pairs (2-4,6-8 and 2-3,4-5) don't overlap, while the remaining four pairs (5-7,7-9, 2-8,3-7, 6-6,4-6, and 2-6,4-8) do overlap:

5-7,7-9 overlaps in a single section, 7.
2-8,3-7 overlaps all of the sections 3 through 7.
6-6,4-6 overlaps in a single section, 6.
2-6,4-8 overlaps in sections 4, 5, and 6.
So, in this example, the number of overlapping assignment pairs is 4.

In how many assignment pairs do the ranges overlap?`

func getInput() string {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

func getAssignments(assgStr string) (int, int) {
	ass := strings.Split(assgStr, "-")
	startIndex, _ := strconv.Atoi(ass[0])
	endIndex, _ := strconv.Atoi(ass[1])
	return startIndex, endIndex
}

func main() {
	input := getInput()
	overlapCount := 0
	for _, line := range strings.Split(input, "\n") {
		pairs := strings.Split(line, ",")
		firstStartIndex, firstEndIndex := getAssignments(pairs[0])
		secondStartIndex, secondEndIndex := getAssignments(pairs[1])
		if (firstStartIndex <= secondStartIndex && secondStartIndex <= firstEndIndex) ||
			(secondStartIndex <= firstStartIndex && firstStartIndex <= secondEndIndex) {
			overlapCount++
		}

	}
	fmt.Println(overlapCount)
}
