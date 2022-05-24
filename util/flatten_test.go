package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Flatten(t *testing.T) {
	jsonInStr := `{
	 "a": 1,
	 "b": true,
	 "c": {
		"d": 3,
		"e": "test"
	 }
	}`

	got := Flatten(jsonInStr)

	want := `{
	  "a": 1,
	  "b": true,
	  "c.d": 3,
	  "c.e": "test"
	}`
	require.JSONEq(t, want, got)
}
