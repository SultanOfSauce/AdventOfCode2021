package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const ReportLength = 12

func main() {
	strs := GatherValues("input.txt")

	vals := make([]uint16, len(strs))

	for i, str := range strs {

		temp, err := strconv.ParseInt(str, 2, 16)

		if err != nil {
			log.Fatal(err)
		}

		vals[i] = uint16(temp)
	}

	//vals = vals[0:100]

	fmt.Println(vals)

	countBits := make([]int, ReportLength)

	for _, val := range vals {
		fmt.Printf("%012b\n", val)
		//fmt.Printf("%d\n", val)
		for i := 0; i < ReportLength; i++ {
			//fmt.Println("Conto:", val, PowTwo(i), val&PowTwo(i))
			//fmt.Println(ReportLength - i + 1)
			countBits[ReportLength-i-1] += GetBit(val, i)
		}
	}

	fmt.Println(countBits, len(vals))

	//DetermineGammaEpsilon(vals, countBits)
	DetermineGases(vals, countBits, true)
	DetermineGases(vals, countBits, false)

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

func DetermineGammaEpsilon(vals []uint16, countBits []int) {

	var noOne, noTwo uint16

	for i, count := range countBits {
		if count >= len(vals)/2 {
			noOne += PowTwo(12 - i)
		}
	}

	noTwo = (^noOne) & 0b0000111111111111

	fmt.Printf("%012b,   %012b\n%d   %d", noOne, noTwo, noOne, noTwo)
}

func DetermineGases(vals []uint16, countBits []int, gas bool) {
	threshold := len(vals) / 2

	for i, count := range countBits {

		firstVal := vals[0]
		over := true

		for _, val := range vals {
			if val != firstVal {
				over = false
				break
			}
		}

		fmt.Println(vals)

		if over {
			fmt.Println("AAAAAAAAAAA:", firstVal, gas)
			return
		}

		newVals := make([]uint16, 0)
		var flag bool
		if gas {
			flag = count >= threshold
		} else {
			flag = count < threshold
		}

		for _, val := range vals {
			if (GetBit(val, ReportLength-i+1) == 1) && flag {
				newVals = append(newVals, val)
			} else if (GetBit(val, ReportLength-i+1) == 0) && !flag {
				newVals = append(newVals, val)
			}
		}
		fmt.Println(newVals)

		vals = newVals
	}
}

func GetBit(number uint16, bit int) int {
	if (number>>bit)&1 == 1 {
		return 1
	} else {
		return 0
	}
}

func PowTwo(exponent int) uint16 {
	var result uint16 = 1
	for i := 1; i < exponent; i++ {
		result *= 2
	}
	return result
}
