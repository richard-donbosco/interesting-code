package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	var output string
	for i := 1; i <= n; i++ {
		if i == 1 {
			output = strconv.Itoa(i)
		} else {
			var current string = ""
			var prev string = ""
			var counter = 1
			var tmp []string = make([]string, 0, 10)
			for j, val := range output {
				current = string(val)
				outputlength := len(output)
				if j == 0 && outputlength == 1 {
					tmp = append(tmp, fmt.Sprintf("%v%v", counter, current))
					break
				}
				if current == prev {
					counter += 1
				} else if current != prev && prev != "" {
					tmp = append(tmp, fmt.Sprintf("%v%v", counter, prev))
					counter = 1
				}
				if outputlength == j+1 {
					tmp = append(tmp, fmt.Sprintf("%v%v", counter, current))
				}
				prev = current
			}
			output = strings.Join(tmp, "")

		}

	}

	return output
}

func main() {
	args := os.Args
	var number int
	if len(os.Args) > 1 {
		number, _ = strconv.Atoi(args[1])
	} else {
		number = 1
	}
	fmt.Printf(countAndSay(number))
}
