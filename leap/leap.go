package leap

const testVersion = 3
// on every year that is evenly divisible by 4
  // except every year that is evenly divisible by 100
  //   unless the year is also evenly divisible by 400



func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
  if (year % 4 == 0) {
    if (year % 100 != 0) {
      return true
    } else if (year % 400 == 0) {
      return true
    }
  }

  return false
}
