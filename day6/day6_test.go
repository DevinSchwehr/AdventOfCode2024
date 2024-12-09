package day6

import "testing"

func TestPartOne(t *testing.T) {
	want := 41
	result, _ := partOne("test.txt")
	if result != want {
		t.Fatalf(`Did not get expected result. Expected %d, received %d`, want, result)
	}
}

func TestPartTwo(t *testing.T) {
	want := 6
	_, guard := partOne("test.txt")
	result := partTwo("test.txt", guard)
	if result != want {
		t.Fatalf(`Did not get expected result. Expected %d, received %d`, want, result)
	}
}
