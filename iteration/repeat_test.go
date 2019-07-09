package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
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
