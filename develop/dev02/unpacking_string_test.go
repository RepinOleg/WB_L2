package dev02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpackingStringCorrect1(t *testing.T) {
	res, err := UnpackingString("a4bc2d5e")
	if assert.Nil(t, err) {
		expectedRes := "aaaabccddddde"
		assert.Equal(t, expectedRes, res, "they should be equal")
	}
}

func TestUnpackingStringCorrect2(t *testing.T) {
	res, err := UnpackingString("abcd")
	if assert.Nil(t, err) {
		expectedRes := "abcd"
		assert.Equal(t, expectedRes, res, "they should be equal")
	}

}

func TestUnpackingStringCorrect3(t *testing.T) {
	res, err := UnpackingString("Г3ж2з")
	if assert.Nil(t, err) {
		expectedRes := "ГГГжжз"
		assert.Equal(t, expectedRes, res, "they should be equal")
	}

}

func TestUnpackingStringCorrectEscape(t *testing.T) {
	res, err := UnpackingString(`qwe\4\5`)
	if assert.Nil(t, err) {
		expectedRes := "qwe45 (*)"
		assert.Equal(t, expectedRes, res, "they should be equal")
	}

}

func TestUnpackingStringCorrectEscape2(t *testing.T) {
	res, err := UnpackingString("qwe\\45")
	if assert.Nil(t, err) {
		expectedRes := "qwe44444 (*)"
		assert.Equal(t, expectedRes, res, "they should be equal")
	}

}

func TestUnpackingStringCorrectEscape3(t *testing.T) {
	res, err := UnpackingString("qwe\\\\5")
	if assert.Nil(t, err) {
		expectedRes := "qwe\\\\\\\\\\ (*)"
		assert.Equal(t, expectedRes, res, "they should be equal")
	}

}

func TestUnpackingStringIncorrect(t *testing.T) {
	res, err := UnpackingString("a422222")
	assert.NotNilf(t, err, err.Error())
	assert.Equal(t, "", res, "they should be equal")
}

func TestUnpackingStringIncorrect1(t *testing.T) {
	res, err := UnpackingString("45")
	assert.NotNilf(t, err, "error message")
	assert.Equal(t, "", res, "they should be equal")
}

func TestUnpackingStringIncorrect2(t *testing.T) {
	res, err := UnpackingString("4")
	assert.NotNilf(t, err, "error message")
	assert.Equal(t, "", res, "they should be equal")
}

func TestUnpackingStringEmpty(t *testing.T) {
	res, err := UnpackingString("")
	assert.Nil(t, err, "error message")
	assert.Equal(t, "", res, "they should be equal")
}
