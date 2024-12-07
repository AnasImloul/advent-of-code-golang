package day_02

func (d Day) firstPart() any {
	res := 0

	for _, report := range readReports() {
		if isReportSafe(report) {
			res++
		}
	}
	return res
}
