package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestUniqueLines(t *testing.T) {
	test := []string{"1", "2", "str", "str", "2"}
	result := UniqueLines(test)
	assert.Equal(t, []string{"1", "2", "str"}, result)
}

func TestIsSorted(t *testing.T) {
	test := []string{"1", "2", "3"}
	result, ok := IsSorted(test)
	assert.Equal(t, true, ok)
	assert.Equal(t, "", result)
}

func TestIsSorted2(t *testing.T) {
	test := []string{"a", "b", "c", "a"}
	result, ok := IsSorted(test)
	assert.Equal(t, false, ok)
	assert.Equal(t, "a", result)
}

func TestSortByColumn(t *testing.T) {
	test := []string{"a 4", "b 3", "c 2", "d 1"}
	f.column = 2
	sort.Slice(test, CompareLines(test))
	expected := []string{"d 1", "c 2", "b 3", "a 4"}
	assert.Equal(t, expected, test)
}

func TestSortByColumn2(t *testing.T) {
	test := []string{"a 4", "b 3", "c 2", "d 1"}
	f.column = 1
	sort.Slice(test, CompareLines(test))
	expected := []string{"a 4", "b 3", "c 2", "d 1"}
	assert.Equal(t, expected, test)
}

func TestSortByNumber(t *testing.T) {
	test := []string{"10", "1", "5", "3"}
	f.numeric = true
	sort.Slice(test, CompareLines(test))
	expected := []string{"1", "3", "5", "10"}
	assert.Equal(t, expected, test)
}

func TestSortByNumberReverse(t *testing.T) {
	test := []string{"10", "1", "5", "3"}
	f.numeric = true
	f.reverse = true
	sort.Slice(test, CompareLines(test))
	expected := []string{"10", "5", "3", "1"}
	assert.Equal(t, expected, test)
}
