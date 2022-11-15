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

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_aAY a as parameter.
 */

func pickingNumbers(a []int) int {
	// Write your code here

	sort.Ints(a)
	max := int(0)
	n := int(len(a))

	start := int(0)

	for end := start + 1; end < n; end++ {
		arr := a[start:end]

		for _, val := range arr {
			if math.Abs(float64(val-a[end])) > 1 {
				currMax := int(len(arr))
				max = maxComparison(&max, &currMax)

				start = end
				break
			}
		}
	}

	lastMax := n - start
	max = maxComparison(&max, &lastMax)
	return max
}

func maxComparison(a, b *int) int {
	if *a > *b {
		return *a
	}
	return *b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
