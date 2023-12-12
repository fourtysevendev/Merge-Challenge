package merge

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var transformErrorToString = cmp.Transformer("transformErrorToString", func(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
})

func Test_ParseInterval(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  error
		expected []Interval
	}{
		{
			name:     "Succesfully parsed interval",
			input:    "[1,10] [2,20] [3,30]",
			wantErr:  nil,
			expected: []Interval{{1, 10}, {2, 20}, {3, 30}},
		},
		{
			name:     "Failure while parsing interval",
			input:    "[1,,10] [2] [3,30]",
			wantErr:  errors.New("inavlid Interval: 1,,10"),
			expected: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := ParseInterval(test.input)
			if !cmp.Equal(actual, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, actual)
			}
			if !cmp.Equal(err, test.wantErr, transformErrorToString) {
				t.Errorf("expected error: %v, got error: %v", test.wantErr, err)
			}
		})
	}
}
