package day_03

import (
	"fmt"
	"regexp"
)

func (d Day) secondPart() {
	input := d.ReadInput()

	pattern := regexp.MustCompile(`(?P<Mul>mul\(\d+,\d+\))|(?P<Do>do\(\))|(?P<Dont>don't\(\))`)

	matches := pattern.FindAllStringSubmatch(input, -1)

	// Extract group names for mapping matches
	groupNames := pattern.SubexpNames()

	var res int64 = 0
	var enabled = true
	for _, match := range matches {

		for i, value := range match {
			// Skip the first element, which is the entire match
			if i == 0 || value == "" {
				continue
			}
			switch groupNames[i] {
			case "Mul":
				if enabled {
					res += d.parseMul(value)
				}
			case "Do":
				enabled = true
			case "Dont":
				enabled = false
			}
		}
	}
	fmt.Println(res)
}
