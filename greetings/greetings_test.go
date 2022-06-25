package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Trent"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Trent")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Trent") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func BenchmarkHelloName(b *testing.B) {
	// run the Hello function b.N times
	for n := 0; n < b.N; n++ {
		Hello("Trent")
	}
}
