package zingsandbox

import "testing"

func TestGreet(t *testing.T) {
	if got, want := Greet("Zing"), "Hello, Zing!"; got != want {
		t.Fatalf("Greet = %q, want %q", got, want)
	}
}
