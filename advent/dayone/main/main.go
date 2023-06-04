package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"util"
)

func getElveArray() []int {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	var arr []int
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var count int
	for fileScanner.Scan() {
		value := fileScanner.Text()
		if value == "" {
			arr = append(arr, count)
			count = 0
		} else {
			convertedValue, _ := strconv.Atoi(value)
			count += convertedValue
		}
	}
	return arr
}

func main() {
	arr := getElveArray()
	maxHeap := util.Heapify(arr)
	value := 0
	topK := 3
	answerOne := heap.Pop(maxHeap).(int)
	for i := 1; i <= topK-1; i++ {
		value += heap.Pop(maxHeap).(int)
	}
	answerPartOne := fmt.Sprintf("Answer for part one: %d", answerOne)
	answerPartTwo := fmt.Sprintf("Answer for part two: %d", answerOne+value)
	fmt.Println(answerPartOne)
	fmt.Println(answerPartTwo)
}
