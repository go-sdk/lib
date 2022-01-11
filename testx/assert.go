package testx

import (
	"github.com/stretchr/testify/assert"
)

var (
	AssertEmpty          = assert.Empty
	AssertNotEmpty       = assert.NotEmpty
	AssertEqual          = assert.Equal
	AssertNotEqual       = assert.NotEqual
	AssertEqualValues    = assert.EqualValues
	AssertNotEqualValues = assert.NotEqualValues
	AssertContains       = assert.Contains
	AssertNotContains    = assert.NotContains
	AssertError          = assert.Error
	AssertNoError        = assert.NoError
)
