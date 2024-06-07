package greetings

import (
	"testing"
	"regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGreetingName(t *testing.T) {
	name := "Damien"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Greeting("Damien")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greeting("Damien") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGreetingEmpty(t *testing.T) {
	msg, err := Greeting("")
	if msg != "" || err == nil {
		t.Fatalf(`Greeting("") = %q, %v, want ""`, msg, err)
	}
}