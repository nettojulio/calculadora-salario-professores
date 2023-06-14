package tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetStringValues(prompt string) string {
	in := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := in.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
	input = strings.Trim(input, " ")
	input = strings.Trim(input, "\n")
	return input
}

func GetFloatValues(prompt string) float64 {
	var (
		input    float64
		strInput string
	)

	fmt.Print(prompt)
	fmt.Scanln(&strInput)
	input, _ = strconv.ParseFloat(strInput, 64)
	return input
}

func GetIntegerValues(prompt string) int64 {
	var (
		input    int64
		strInput string
	)
	fmt.Print(prompt)
	fmt.Scanln(&strInput)
	input, _ = strconv.ParseInt(strInput, 0, 64)
	return input
}
