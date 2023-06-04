package testing

import "testing"

func TestBuildJoin(t *testing.T) {
	//t.Error("no test written yet")
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}

	if joinWithCommas(list) != "apple and orange" {
		t.Error("Didn't match expected value")
	}

	//t.Error("no test written yet")
}

func TestThreeElements(t *testing.T) {
	//t.Error("no test written yet")

	list := []string{"apple", "orange", "peach"}

	if joinWithCommas(list) != "apple, orange, and peach" {
		t.Error("Didn't match expected value")
	}
}
