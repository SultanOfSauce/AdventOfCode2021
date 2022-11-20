package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const WindowSize = 3

func main() {

	vals := GatherValues("input.txt")

	windowedValues := make([]int, len(vals)-WindowSize)

	counter := 1

	for i := range windowedValues {
		sum := 0
		for j := 0; j < WindowSize; j++ {
			sum += vals[i+j]
		}
		windowedValues[i] = sum
	}

	for i := 1; i < len(windowedValues); i++ {
		if windowedValues[i]-windowedValues[i-1] > 0 {
			counter++
		}
	}

	fmt.Println(counter)
}

func GatherValues(filename string) (vals []int) {
	bytes, err := os.ReadFile(filename)
	content := string(bytes)

	if err != nil {
		log.Fatal(err)
	}

	content = strings.TrimSpace(content)
	strs := strings.Split(string(content), "\n")

	vals = make([]int, len(strs))

	for i := range vals {
		vals[i], _ = strconv.Atoi(strs[i])
	}

	return
}
