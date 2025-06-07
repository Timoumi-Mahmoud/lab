package iteration

import "testing"

func TestRepeat(t *testing.T) {
	// t.Fatal("not implemented")
	testCases := []struct {
		repeated    string
		repeatCount int
		expected    string
	}{
		{"a", 5, "aaaaa"},
		{"hello", 2, "hellohello"},
		{"", 0, ""},
	}

	for _, v := range testCases {
		result := Repeat(v.repeated, v.repeatCount)
		if result != v.expected {
			t.Errorf("expected %q but got %q", v.expected, result)
		}

	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}

}

/*
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/Timoumi-Mahmoud/GOLANG/21-testing/Go-TDD/iteration
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkRepeat-8   	 9028662(nbre of times)	       124.7 ns/op (func takes on average 124.7 ns )
PASS
ok  	github.com/Timoumi-Mahmoud/GOLANG/21-testing/Go-TDD/iteration	1.265s

--------Using strings.Builder and WriteString-------
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkRepeat-8   	34390213	        33.67 ns/op
PASS
o
*/
