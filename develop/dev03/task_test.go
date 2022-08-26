package main

import "testing"

func TestSortByColumn(t *testing.T) {
	input := []string{"some\ntest\nstring\nwith\nfew\nlines"}
	expectations := []string{"few\nlines\nsome\nstring\ntest\nwith"}

	for i := 0; i < len(input) && i < len(expectations); i++ {
		result := sortByColumn(input[i], 0)

		if result != expectations[i] {
			t.Fatalf("\nbad result for %s\nexpected %s\ngot %s\n",
				input[i], expectations[i], result)
		}
	}
}

func TestSortByNumbers(t *testing.T) {
	input := []string{"some\n2\ntest\n1\nstring\n10\n3\n6\nwith\n7\nfew\nlines"}
	expectations := []string{"1\n2\n3\n6\n7\n10\nsome\ntest\nstring\nwith\nfew\nlines"}

	for i := 0; i < len(input) && i < len(expectations); i++ {
		result := sortByNumbers(input[i])

		if result != expectations[i] {
			t.Fatalf("\nbad result for %s\nexpected %s\ngot %s\n",
				input[i], expectations[i], result)
		}
	}
}

func TestReverse(t *testing.T) {
	input := []string{"some\ntest\nstring\nwith\nfew\nlines"}
	expectations := []string{"with\ntest\nstring\nsome\nlines\nfew"}

	for i := 0; i < len(input) && i < len(expectations); i++ {
		result := reverse(input[i])

		if result != expectations[i] {
			t.Fatalf("\nbad result for %s\nexpected %s\ngot %s\n",
				input[i], expectations[i], result)
		}
	}
}

func TestSelectUnic(t *testing.T) {
	input := []string{"some\ntest\ntest\nstring\nwith\ntest\nfew\nlines\nfew"}
	expectations := []string{"some\ntest\nstring\nwith\nfew\nlines"}

	for i := 0; i < len(input) && i < len(expectations); i++ {
		result := selectUnic(input[i])

		if result != expectations[i] {
			t.Fatalf("\nbad result for %s\nexpected %s\ngot %s\n",
				input[i], expectations[i], result)
		}
	}
}
