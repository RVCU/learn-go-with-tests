package iteration

import "testing"

func TestRepeat(t *testing.T) {
    repeated := Repeat("racecar")
    expected := "racecarracecarracecarracecarracecar"
    if repeated != expected {
        t.Errorf("expected, '%s' but got '%s'", expected, repeated)
    }
}

func Repeat(daString string) string {
    var repeatedString string
    for i := 0; i < 5; i++ {
       repeatedString += daString
    }
    return repeatedString
}
