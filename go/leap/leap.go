// Package leap contains functions for working with leap years
package leap

// IsLeapYear returns true if provided year is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
