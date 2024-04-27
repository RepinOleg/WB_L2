package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	test := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	res := GroupAnagrams(test)
	expected := map[string][]string{
		"листок": {"листок", "слиток", "столик"},
		"пятак":  {"пятак", "пятка", "тяпка"},
	}

	assert.Equal(t, expected, res)
}

func TestGroupAnagrams2(t *testing.T) {
	test := []string{"пятак"}
	res := GroupAnagrams(test)
	expected := map[string][]string{}

	assert.Equal(t, expected, res)
}

func TestGroupAnagrams3(t *testing.T) {
	test := []string{"пятак", "пятак", "пятак", "пятак", "тяпка"}
	res := GroupAnagrams(test)
	expected := map[string][]string{
		"пятак": {"пятак", "тяпка"},
	}

	assert.Equal(t, expected, res)
}
