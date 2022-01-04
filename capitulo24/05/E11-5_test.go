package main

import (
	"testing"
)

func TestSum1_2_3_4(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	sum := SumAll(numbers...)
	if sum != 10 {
		t.Errorf("When sum 1, 2, 3 and 4 should returns 10 but returned %d", sum)
	}
}

func TestSum22_33(t *testing.T) {
	numbers := []int{22, 33}
	sum := SumAll(numbers...)
	if sum != 55 {
		t.Errorf("When sum 22 and 33 should returns 55 but returned %d", sum)
	}
}
