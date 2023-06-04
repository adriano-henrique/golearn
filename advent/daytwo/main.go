package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	partOneSolution()
	partTwoSolution()
}

func partOneSolution() {
	equivalentMapping := getEquivalentMapping()

	playScoreMapping := getPlayScoreMapping()

	playBeatMapping := getPlayBeatMapping()

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var count int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitLine := strings.Fields(line)
		oppPlay := splitLine[0]
		yourPlay := splitLine[1]
		equivalentOppPlay := equivalentMapping[oppPlay]
		count += playScoreMapping[yourPlay]
		if playBeatMapping[equivalentOppPlay] == yourPlay {
			count += 6
		} else if equivalentOppPlay == yourPlay {
			count += 3
		}
	}
	fmt.Println(count)
}

func partTwoSolution() {
	equivalentMapping := getEquivalentMapping()

	playBeatMapping := getPlayBeatMapping()

	playScoreMapping := getPlayScoreMapping()

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var count int
	var additionalCount int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitLine := strings.Fields(line)
		oppPlay := splitLine[0]
		shouldPlay := splitLine[1]
		equivalenteOppPlay := equivalentMapping[oppPlay]
		if shouldPlay == "Y" {
			additionalCount = playScoreMapping[equivalenteOppPlay] + 3
		} else if shouldPlay == "X" {
			reversedPlayBeatMapping := reverseMap(playBeatMapping)
			yourMove := reversedPlayBeatMapping[equivalenteOppPlay]
			additionalCount = playScoreMapping[yourMove]
		} else {
			yourMove := playBeatMapping[equivalenteOppPlay]
			additionalCount = playScoreMapping[yourMove] + 6
		}
		count += additionalCount
	}
	fmt.Println(count)
}

func getEquivalentMapping() map[string]string {
	return map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
}

func getPlayScoreMapping() map[string]int {
	return map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
}

func getPlayBeatMapping() map[string]string {
	return map[string]string{
		"X": "Y",
		"Y": "Z",
		"Z": "X",
	}
}

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}
