package clock

import "fmt"

type Clock struct {
  hour int
  minute int
}

func New(hour, minute int) Clock {
  return Clock{hour, minute}.calculate()
}

func (c Clock) calculate() Clock {

  if c.minute < 0 {
    c.hour -= -1 * c.minute / 60 + 1
    c.minute = 60 - (-1 * c.minute % 60)
  }

  c.hour += c.minute / 60
  if c.minute / 60 > 0 {
    c.minute -= c.minute / 60 * 60
  } else {
    c.minute = c.minute % 60
  }

  if c.hour < 0 {
    c.hour = 24 - (-1 * c.hour % 24)
  }

  c.hour = c.hour % 24

  return Clock{hour: c.hour, minute: c.minute}
}

func (c Clock) Add(add int) Clock {
  return New(c.hour, c.minute + add)
}

func (c Clock) Subtract(add int) Clock {
  return New(c.hour, c.minute - add)
}

func (c Clock) String() string {
  return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}