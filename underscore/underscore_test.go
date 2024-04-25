package underscore

import "testing"

func TestCamel(t *testing.T) {
	testCases := [...]struct {
		arg  string
		want string
	}{
		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
		{"with a space", "with a space"},
		{"endsWithA", "ends_with_a"},
	}
	for _, tt := range testCases {
		t.Logf("Testing: %q", tt.arg)
		if got := Camel(tt.arg); got != tt.want {
			t.Errorf("Camel(%q) = %q, want %q", tt.arg, got, tt.want)
		}
	}
}
