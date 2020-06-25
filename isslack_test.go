package isslack

import "testing"

func TestIsslack(t *testing.T) {
	want := "200 OK"
	if got := Isslack("Test", "https://hooks.slack.com/services/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"); got != want {
		t.Errorf("Isslack() = %q, want %q", got, want)
	}
}
