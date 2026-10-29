// A palindromic number reads the same both ways. The smallest 6 digit palindrome made from the product of
// two 3-digit numbers is 101101 = 143 * 707.

// Find the largest palindrome made from the product of two 3-digit numbers which is less than N.

// Input Format:
// First line contains T that denotes the number of test cases. This is followed by T lines, each containing an integer, N.

// Constraints:
// - 1 <= T <= 100
// - 101101 < N < 1000000

// Output Format:
// Print the required answer for each test case in a new line.

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func isProductOfThreeDigitNumbers(number int) bool {
	lowerBound, upperBound := max(100, (number+998)/999), int(math.Sqrt(float64(number)))

	for factor := lowerBound; factor <= upperBound; factor++ {
		if number%factor == 0 {
			return true
		}
	}

	return false
}

func generateValidPalindromes() []int {
	validPalindromes := make([]int, 0, 450)

	for firstExtreme := 1; firstExtreme < 10; firstExtreme++ {
		for secondExtreme := range 10 {
			for thirdExtreme := range 10 {
				number := (100000+1)*firstExtreme + (10000+10)*secondExtreme + (1000+100)*thirdExtreme

				if isProductOfThreeDigitNumbers(number) {
					validPalindromes = append(validPalindromes, number)
				}
			}
		}
	}

	return validPalindromes
}

func main() {
	validPalindromes := generateValidPalindromes()
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	for range tTemp {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		fmt.Println(validPalindromes[sort.SearchInts(validPalindromes, int(nTemp-1))-1])
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
