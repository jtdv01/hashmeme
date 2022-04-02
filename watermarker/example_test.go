package watermarker

import (
    "testing"
)

func TestImport(t *testing.T) {
    got := Foo()
    want := "fox1o"

    if got != want {
        t.Errorf("Got %q, want: %q", got, want)
    }
}