// Clock stub file

package clock

import (
  "fmt"
  "strconv"
)

const testVersion = 4


type Clock struct {
  hour int
  minute int
}

func New(hour, minute int) Clock {

  var hrs, min int
  extra_hours := minute / 60
  if (minute < 0) {
    leftover := minute % 60
    if (leftover != 0) {
        extra_hours -= 1
    }
    min = 60 + leftover
  } else {
    min = minute % 60
  }

  if tot_hrs := hour + extra_hours; tot_hrs < 0 {
    hrs = 24 + tot_hrs % 24
  } else {
    hrs = tot_hrs % 24
  }

  var clock Clock
  clock.hour = hrs
  clock.minute = min
  return clock
}

func (c Clock) String() string {
  if c.minute > 59 {
		c.hour = c.hour + c.minute/60
	} else if c.minute < 0 {
		c.hour = c.minute/60 - 1
	}

	c.minute = c.minute % 60
	c.hour = c.hour % 24

	min := strconv.Itoa(c.minute)
	hour := strconv.Itoa(c.hour)
	if c.minute < 10 && c.minute > -1 {
		min = "0" + min
	}

	if c.hour < 10 && c.hour > -1 {
		hour = "0" + hour
	}

	return hour + ":" + min
}

func stringify(num int) string {
  var res string
  if (num < 10) {
    res = "0" + fmt.Sprintf("%v", num)
  } else {
    res = fmt.Sprintf("%v", num)
  }
  return res
}

func (c Clock) Add(minutes int) Clock {
  c.minute = minutes + c.minute
  	if c.minute > -1 && c.minute < 60 {
  		return c
  	}

  	plusHour := int(c.minute / 60)
  	c.hour = plusHour + c.hour
  	c.minute = c.minute % 60
  	if c.minute < 0 {
  		c.hour--
  		c.minute = c.minute + 60
  	}
  	if c.hour < 0 {
  		c.hour = c.hour % 24
  		c.hour = c.hour + 24
  	}

  	return c
  }
