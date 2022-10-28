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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	lH := len(s)

	//sString := s[:2]  // 12
	amPm := s[lH-2:]      // :00:00
	middle := s[2 : lH-2] // AM

	//fmt.Print(sString, middle, amPm)
	//fmt.Printf("\n %v - %T", sString, sString)

	sInt, err := strconv.Atoi(s[:2])
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("\n %v - %T", sInt, sInt) ; fmt.Println()

	if sInt == 12 && amPm == "AM" {
		sInt = 0
	}
	if sInt != 12 && amPm == "PM" {
		sInt += 12
	}

	//sStr := strconv.Itoa(sInt)
	//fmt.Print(sStr,middle)
	return fmt.Sprintf("%02d%s", sInt, middle)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
