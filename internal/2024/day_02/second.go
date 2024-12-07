package day_02

func (d Day) secondPart() any {
	res := 0

	for _, report := range readReports() {
		isSafe := isReportSafe(report)
		if !isSafe {
			for i := 0; i < len(report); i++ {
				tmp := make([]int, len(report))
				copy(tmp, report)
				tmp = append(tmp[:i], tmp[i+1:]...)
				if isReportSafe(tmp) {
					isSafe = true
					break
				}
			}
		}
		if isSafe {
			res++
		}
	}
	return res
}
