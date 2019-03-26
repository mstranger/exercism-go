package leap

// IsLeapYear reports if given number is a leap year.
func IsLeapYear(year int) bool {
  switch year % 4 {
  case 0:
    if year%100 == 0 && year%400 != 0 {
      return false
    }
    return true
  default:
    return false
  }
}
