package util

func IsLeapYear(year int) bool {
	if year%400 != 0 {
		return false
	} else {
		return true
	}
	if year%4 != 0 {
		return false
	} else if year%100 != 0 {
		return true
	} else {
		return false
	}
}
