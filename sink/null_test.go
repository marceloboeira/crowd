package sink

import (
	"strings"
	"testing"
)

func TestNoError(t *testing.T) {
	q := NewNull()

	err := q.Push(strings.NewReader("this should not error"))

	if err != nil {
		t.Error("", err)
	}
}
