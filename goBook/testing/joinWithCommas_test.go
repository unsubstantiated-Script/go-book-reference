package testing

import (
	"fmt"
	"testing"
)

func TestBuildJoin(t *testing.T) {
	//t.Error("no test written yet")
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}

	want := "apple and orange"
	got := joinWithCommas(list)

	if got != want {
		t.Error(errorString(list, got, want))
	}

	//t.Error("no test written yet")
}

func TestThreeElements(t *testing.T) {
	//t.Error("no test written yet")

	list := []string{"apple", "orange", "peach"}

	want := "apple, orange, and peach"
	got := joinWithCommas(list)

	if got != want {
		t.Error(errorString(list, got, want))
	}

}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("\n got: joinWithCommas(%#v) = \"%s\"\n want: \"%s\"", list, got, want)
}
