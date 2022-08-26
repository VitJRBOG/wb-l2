package main

import "testing"

func TestUnpack(t *testing.T) {
	input := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}
	expectations := []string{"aaaabccddddde", "abcd", "", "", "qwe45", "qwe44444", "qwe\\\\\\\\\\"}

	for i := 0; i < len(input) && i < len(expectations); i++ {
		result, err := unpack(input[i])
		if err != nil {
			if input[i] != "45" || err.Error() != "error: invalid line" {
				t.Fatalf("\ngot error: %s\nlast input %s\n",
					err.Error(), input[i])
			}
		}

		if result != expectations[i] {
			t.Fatalf("\nbad result for %s\nexpected %s\ngot %s\n",
				input[i], expectations[i], result)
		}
	}
}
