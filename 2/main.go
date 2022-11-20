package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	strs := GatherValues("input.txt")

	var instr []string
	var hor, depth, aim int

	for _, movement := range strs {

		instr = strings.Split(movement, " ")

		val, _ := strconv.Atoi(instr[1])
		switch instr[0] {
		case "forward":
			hor += val
			depth += aim * val
		case "up":
			aim -= val
		case "down":
			aim += val
		}

	}

	fmt.Println(hor, " ", depth, " ", hor*depth)
}

func GatherValues(filename string) (strs []string) {
	bytes, err := os.ReadFile(filename)
	content := string(bytes)

	if err != nil {
		log.Fatal(err)
	}

	content = strings.TrimSpace(content)
	strs = strings.Split(string(content), "\n")

	return
}
