package day_02

import "fmt"

func (d Day) firstPart() {
	res := 0

	for _, report := range readReports() {
		if isReportSafe(report) {
			res++
		}
	}
	fmt.Println(res)
}
