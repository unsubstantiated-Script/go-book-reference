package testing

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestBuildJoin(t *testing.T) {
	tests := []testData{
		{list: []string{}, want: ""},
		{list: []string{"apple"}, want: "apple"},
		{list: []string{"apple", "orange"}, want: "apple and orange"},
		{list: []string{"apple", "orange", "peach"}, want: "apple, orange, and peach"},
	}

	for _, test := range tests {
		got := joinWithCommas(test.list)
		if got != test.want {
			t.Errorf("\n got: joinWithCommas(%#v) = \"%s\" \n want: \"%s\"", test.list, got, test.want)
		}
	}
	//t.Error("no test written yet")
}

//func TestOneElement(t *testing.T) {
//	list := []string{"apple"}
//
//	want := "apple"
//	got := joinWithCommas(list)
//
//	if got != want {
//		t.Error(errorString(list, got, want))
//	}
//
//	//t.Error("no test written yet")
//}
//
//func TestTwoElements(t *testing.T) {
//	list := []string{"apple", "orange"}
//
//	want := "apple and orange"
//	got := joinWithCommas(list)
//
//	if got != want {
//		t.Error(errorString(list, got, want))
//	}
//
//	//t.Error("no test written yet")
//}
//
//func TestThreeElements(t *testing.T) {
//	//t.Error("no test written yet")
//
//	list := []string{"apple", "orange", "peach"}
//
//	want := "apple, orange, and peach"
//	got := joinWithCommas(list)
//
//	if got != want {
//		t.Error(errorString(list, got, want))
//	}
//
//}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("\n got: joinWithCommas(%#v) = \"%s\"\n want: \"%s\"", list, got, want)
}
