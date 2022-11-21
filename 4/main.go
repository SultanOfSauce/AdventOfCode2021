package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoSheet [5][5]int

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	numsString := scanner.Text()
	var nums []int

	{
		numsSlice := strings.Split(numsString, ",")
		nums = make([]int, len(numsSlice))

		for i, num := range numsSlice {
			nums[i], _ = strconv.Atoi(num)
		}
	}

	sheets := make([]BingoSheet, 0)

	{
		var tempMatrix BingoSheet = NewBingoSheet()

		for scanner.Scan() {
			txt := scanner.Text()
			if txt == "\n" {
				continue
			}

		}
	}
}

func ParseRow(txt string) [5]int {

}

func NewBingoSheet() BingoSheet {

	var b BingoSheet

	for i := 0; i < 5; i++ {
		b[i] = [5]int{}
	}

	return b
}
