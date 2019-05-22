package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
// This passes 8 out of 12 test cases
func minimumBribes(q []int32) {
	minBribes := 0
	chaos := false

	for key, value := range q {
		v := int(value)

		if v > (key + 3) {
			chaos = true
			break
		}

		// count the number of people who have bribed their way ahead of you
		for i := 0; i <= key; i++ {
			if int(q[i]) > v {
				minBribes++
			}
		}
	}
	if chaos == true {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(minBribes)
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
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
