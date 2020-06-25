package isslack

import "testing"

// TODO:
// must take SLACK URL as argument
func TestIsslack(t *testing.T) {
	want := "200 OK"
	if got := Isslack("Test", "SLACK URL"); got != want {
		t.Errorf("Isslack() = %q, want %q", got, want)
	}
}
