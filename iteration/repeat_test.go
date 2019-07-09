package iteration

import "fmt"
import "testing"

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ { //https://golang.org/pkg/testing/#hdr-Benchmarks
        Repeat("a",1000000000)
    }
}
func ExampleRepeat() {
    returnString := Repeat("b",3)
    fmt.Println(returnString)
    //Output: bbb
}
func TestRepeat(t *testing.T) {
    repeated := Repeat("racecar",5)
    expected := "racecarracecarracecarracecarracecar"
    if repeated != expected {
        t.Errorf("expected, '%s' but got '%s'", expected, repeated)
    }
}

func Repeat(daString string, repeatNumber int) string {
    var repeatedString string
    for i := 0; i < repeatNumber; i++ {
       repeatedString += daString
    }
    return repeatedString
}
