package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var counter int = 1
	sc := bufio.NewScanner(os.Stdin)

	// regex for matching timestamp lines
	var timing_regex string = "([\\d:\\.]{9,12})\\s*-->\\s*([\\d:\\.]{9,12})"
	re := regexp.MustCompile(timing_regex)

	// regex for matching tags
	tag_re := regexp.MustCompile("<.*>")

	for sc.Scan() {
		var line string = sc.Text()
		line = strings.TrimSpace(line)

		// check if the current line contains timestamps
		if re.MatchString(line) {
			// print and increment the counter
			fmt.Println(counter)
			counter++

			// match for just the timestamps to filter out any additional
			// webvtt format options
			match := re.FindString(line)
			fmt.Println(strings.ReplaceAll(match, ".", ","))

			for sc.Scan() {
				line = sc.Text()
				line = tag_re.ReplaceAllString(line, "")
				var result string
				if len(line) == 0 {
					fmt.Println(result)
					break
				}
				result += line
				fmt.Println(result)
			}
		}
	}
}
