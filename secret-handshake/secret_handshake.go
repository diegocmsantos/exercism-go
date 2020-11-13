package secret

var (
  moves = []struct {
    i    uint
    move string
  }{
    {1, "wink"},
    {2, "double blink"},
    {4, "close your eyes"},
    {8, "jump"},
  }
)

// Handshake turns a number into a secret handshake sequence for those in the
// know
func Handshake(i uint) []string {
  var hs []string

  if i < 0 {
    return nil
  }

  for _, move := range moves {
    if i&move.i == move.i {
      hs = append(hs, move.move)
    }
  }

  if i&16 == 16 {
    reverse(hs)
  }

  return hs
}

// Lifted from https://github.com/golang/go/wiki/SliceTricks.
// I'm not happy about this.
func reverse(s []string) {
  for i := len(s)/2 - 1; i >= 0; i-- {
    opp := len(s) - 1 - i
    s[i], s[opp] = s[opp], s[i]
  }
}
