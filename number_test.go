package jsonwat_test

import (
	"encoding/json"
	"testing"

	"github.com/imperfectgo/jsonwat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func toIntSlice(src []jsonwat.Int) []int {
	dst := make([]int, len(src))
	for i, v := range src {
		dst[i] = int(v)
	}
	return dst
}

func TestInt_UnmarshalJSON_Simple(t *testing.T) {
	data := []byte(`{
	"int": 1234,
	"str": "1234",
	"int_arr": [1, 2, 3],
	"str_arr": ["1", "2", "3"],
	"mixed_arr": ["1", 2, "3"]
}`)

	var obj struct {
		Int      jsonwat.Int   `json:"int"`
		Str      jsonwat.Int   `json:"str"`
		IntArr   []jsonwat.Int `json:"int_arr"`
		StrArr   []jsonwat.Int `json:"str_arr"`
		MixedArr []jsonwat.Int `json:"mixed_arr"`
	}
	err := json.Unmarshal(data, &obj)
	require.NoError(t, err)

	assert.EqualValues(t, 1234, obj.Int)
	assert.EqualValues(t, 1234, obj.Str)
	assert.Equal(t, []int{1, 2, 3}, toIntSlice(obj.IntArr))
	assert.Equal(t, []int{1, 2, 3}, toIntSlice(obj.StrArr))
	assert.Equal(t, []int{1, 2, 3}, toIntSlice(obj.MixedArr))
}

func TestInt_UnmarshalJSON_Error(t *testing.T) {
	fixtures := []string{
		// array
		`{"number": []}`,
		// object
		`{"number": {}}`,
		// null
		`{"number": null}`,
		// boolean
		`{"number": true}`,
		`{"number": false}`,
		// invalid integer
		`{"number": 123.45}`,
		// invalid integer string
		`{"number": "123.45"}`,
	}

	var obj struct {
		Number jsonwat.Int `json:"number"`
	}
	for i, fixture := range fixtures {
		err := json.Unmarshal([]byte(fixture), &obj)
		assert.Errorf(t, err, "fixture #%d: %v", i, fixture)
	}
}
