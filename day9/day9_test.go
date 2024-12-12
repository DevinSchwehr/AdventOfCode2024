package day9

import "testing"

func TestPartOne(t *testing.T) {
	want := 1928
	result := partOne("test.txt")
	if result != want {
		t.Fatalf(`Did not get expected result. Expected %d, received %d`, want, result)
	}
}

func TestPartTwo(t *testing.T) {
	want := 2858
	result := partTwo("input.txt")
	if result != want {
		t.Fatalf(`Did not get expected result. Expected %d, received %d`, want, result)
	}
}
