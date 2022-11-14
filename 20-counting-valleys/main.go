package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int, path string) (valleyCounter int) {
	pathList := strings.Split(path, "")
	//fmt.Printf("pathlist: %s", pathList)
	currentElevation := 0
	var enteredValley bool

	for i := 0; i < steps; i++ {
		if pathList[i] == "U" {
			currentElevation++
		} else {
			currentElevation--
		}
		if currentElevation < 0 && enteredValley == false {
			enteredValley = true
			valleyCounter++
		}
		if currentElevation >= 0 && enteredValley == true {
			enteredValley = false
		}
	}
	return valleyCounter
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	steps := int(stepsTemp)

	path := readLine(reader)

	result := countingValleys(steps, path)

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
