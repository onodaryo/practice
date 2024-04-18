package util

func IsLeapYear(year int) bool {
	var check bool = false
	if year%4 == 0 {
		if year%100 != 0 {
			check = true
		}
	}
	if year%400 == 0 {
		check = true
	}
	return check
}
