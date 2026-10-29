// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible(divisible with no remainder) by all of the numbers from 1 to N?

// Input Format:
// First line contains T that denotes the number of test cases. This is followed by T lines, each containing an integer, N.

// Constraints:
// - 1 <= T <= 10
// - 1 <= N <= 40

// Output Format:
// Print the required answer for each test case.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func hcf(a, b int) int {
	if b < a {
		return hcf(b, a)
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func generateCumulativeLcm() [41]int {
	cumulativeLcm := [41]int{}
	cumulativeLcm[1] = 1

	for number := 2; number <= 40; number++ {
		prevLcm := cumulativeLcm[number-1]
		cumulativeLcm[number] = (prevLcm / hcf(prevLcm, number)) * number
	}

	return cumulativeLcm
}

func main() {
	cumulativeLcm := generateCumulativeLcm()
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	for range tTemp {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		fmt.Println(cumulativeLcm[nTemp])
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
