package main

import (
	"aoc2021/day1"
	"fmt"
	"io/ioutil"
	"os"
)

func loadContent(fileName string) string {
	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(fileBytes)
}

func solveDay1() {
	content := loadContent("input/day1.txt")
	part1Result, err := day1.Part1(content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Day 1 Part 1 %d\n", part1Result)
}

func main() {
	solveDay1()
}
