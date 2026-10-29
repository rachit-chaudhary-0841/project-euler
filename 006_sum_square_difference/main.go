// The sum of the squares of the first ten natural numbers is, 1^2 + 2^2 + ... + 10^2 = 385. The square of the sum of
// the first ten natural numbers is, (1 + 2 + ... + 10)^2 = 55^2 = 3025. Hence the absolute difference between the sum
// of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.

// Find the absolute difference between the sum of the squares of the first N natural numbers and the square of the sum.

// Input Format:
// First line contains T that denotes the number of test cases. This is followed by T lines, each containing an integer, N.

// Constraints:
// - 1 <= T <= 10^4
// - 1 <= N <= 10^4

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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	for range tTemp {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int(nTemp)

		fmt.Println(n * (n + 1) * ((3 * n * (n + 1)) - (4 * n) - 2) / 12)
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
