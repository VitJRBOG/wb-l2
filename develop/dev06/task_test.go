package main

import (
	"testing"
)

func TestSelectColumn(t *testing.T) {
	input1 := []string{
		"ColumnOne	ColumnTwo	ColumnThree	ColumnFour",
		"ColumnOne ColumnTwo ColumnThree ColumnFour",
		"ColumnOne|ColumnTwo|ColumnThree|ColumnFour",
	}

	input2 := []string{
		"	",
		" ",
		" ",
	}

	input3 := []bool{
		false,
		false,
		true,
	}

	input4 := []int{
		3,
		3,
		0,
	}

	expectations := []string{
		"ColumnFour",
		"ColumnFour",
		"",
	}

	for i := 0; i < len(input1); i++ {
		result := SelectColumn(input1[i], input2[i], input3[i], input4[i])
		if expectations[i] != result {
			t.Fatalf("\nbad result for: \n\"%s\"\n\"%s\"\n\"%v\"\n\"%d\"\nexpected: \n\"%s\"\ngot: \n\"%s\"\n",
				input1[i], input2[i], input3[i], input4[i], expectations[i], result)
		}
	}
}
