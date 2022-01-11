package testx

import (
	"github.com/stretchr/testify/require"
)

var (
	RequireEmpty          = require.Empty
	RequireNotEmpty       = require.NotEmpty
	RequireEqual          = require.Equal
	RequireNotEqual       = require.NotEqual
	RequireEqualValues    = require.EqualValues
	RequireNotEqualValues = require.NotEqualValues
	RequireContains       = require.Contains
	RequireNotContains    = require.NotContains
	RequireError          = require.Error
	RequireNoError        = require.NoError
)
