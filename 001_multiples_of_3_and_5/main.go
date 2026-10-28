// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below N.

// Input Format:
// First line contains T that denotes the number of test cases. This is followed by T lines, each containing an integer, N.

// Constraints:
// - 1 <= T <= 10^5
// - 1 <= N <= 10^9

// Output Format:
// For each test case, print an integer that denotes the sum of all the multiples of 3 or 5 below N.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	for range tTemp {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int(nTemp - 1)

		if n < 3 {
			fmt.Println(0)
			continue
		}
		multipleSum := (n / 3) * (3 + n - (n % 3)) / 2

		if n < 5 {
			fmt.Println(multipleSum)
			continue
		}
		multipleSum += (n / 5) * (5 + n - (n % 5)) / 2

		if n < 15 {
			fmt.Println(multipleSum)
			continue
		}
		multipleSum -= (n / 15) * (15 + n - (n % 15)) / 2

		fmt.Println(multipleSum)
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
