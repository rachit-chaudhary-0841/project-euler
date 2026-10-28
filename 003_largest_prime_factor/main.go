// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of a given number N?

// Input Format:
// First line contains T, the number of test cases. This is followed by T lines each containing an integer N.

// Constraints:
// - 1 <= T <= 10
// - 10 <= N <= 10^12

// Output Format:
// For each test case, display the largest prime factor of N.

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
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		var largestPrimeFactor int64 = 1

		for n%2 == 0 {
			largestPrimeFactor = 2
			n /= 2
		}

		for n%3 == 0 {
			largestPrimeFactor = 3
			n /= 3
		}

		for factor := int64(5); factor*factor <= n; factor += 6 {
			for n%factor == 0 {
				largestPrimeFactor = factor
				n /= factor
			}

			for n%(factor+2) == 0 {
				largestPrimeFactor = factor + 2
				n /= (factor + 2)
			}
		}

		if n > 1 {
			largestPrimeFactor = n
		}

		fmt.Println(largestPrimeFactor)
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
