package day2

import "testing"

func TestPartOne(t *testing.T) {
	want := 2
	result := partOne("test.txt")
	if result != want {
		t.Fatalf(`Did not get expected result. Expected %d, received %d`, want, result)
	}
}

func TestPartTwo(t *testing.T) {
	want := 9
	result := partTwo("test.txt")
	if result != want {
		t.Fatalf(`Did not get expected result. Expected %d, received %d`, want, result)
	}
}
