package day_03

import (
	"fmt"
	"regexp"
)

func (d Day) firstPart() any {
	input := d.ReadInput()

	pattern := regexp.MustCompile("mul\\(\\d+,\\d+\\)")

	fmt.Println(pattern)

	matches := pattern.FindAllStringSubmatch(input, -1)

	var res int64 = 0
	for _, m := range matches {
		res += d.parseMul(m[0])
	}
	return res
}
