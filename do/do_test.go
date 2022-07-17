package do

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFull(t *testing.T) {
	tests := map[string]struct {
		someString string
		someInt    int
		someBool   bool
		exp        string
		expErr     string
	}{
		"1": {someString: "a", someInt: 1, someBool: true, exp: "[1A]", expErr: ""},
		"2": {someString: "a", someInt: 13, someBool: true, exp: "A", expErr: ""},
		"3": {someString: "", someInt: 1, someBool: true, exp: "", expErr: "invalid s"},
		"4": {someString: "", someInt: 0, someBool: true, exp: "", expErr: "invalid s"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Do(tt.someString, tt.someInt, tt.someBool)
			assert.Equal(t, tt.exp, actual)
			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
