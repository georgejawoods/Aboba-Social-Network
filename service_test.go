package aboba

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_isID(t *testing.T) {
	tt := []struct {
		in   string
		want bool
	}{
		{"", false},
		{"cqefe7lo37026c7q387g", true},
	}
	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			got := isID(tc.in)
			assert.Equal(t, tc.want, got, "isID(%q) = %v, want %v", tc.in, got, tc.want)
		})
	}
}
