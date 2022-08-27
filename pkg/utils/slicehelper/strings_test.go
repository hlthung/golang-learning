package slicehelper

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateStrings(t *testing.T) {
	in := []string{"s1", "s2", "s3", "s1", "s1", "s2"}
	expected := []string{"s1", "s2", "s3"}

	deduped, dupes := RemoveDuplicateStrings(in)
	assert.Equal(t, expected, deduped)
	assert.Equal(t, []string{"s1", "s1", "s2"}, dupes)
}

func TestContainsString(t *testing.T) {
	testData := []struct {
		a []string
		b string
		c bool
	}{
		{[]string{"a", "b", "c"}, "b", true},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{"a", "b", "cDe"}, "cde", false},
		{[]string{"dsadas", "2313", "vcx"}, "dsadas", true},
	}

	for _, td := range testData {
		output := ContainsString(td.a, td.b)
		assert.Equal(t, td.c, output, "Wrong output for contains string!", td.a, td.b, td.c)
	}
}

func TestContainsStringInsensitive(t *testing.T) {
	testData := []struct {
		a []string
		b string
		c bool
	}{
		{[]string{"a", "b", "c"}, "b", true},
		{[]string{"a", "b", "cDe"}, "cde", true},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{"dsadas", "2313", "vcx"}, "dsadas", true},
	}

	for _, td := range testData {
		output := ContainsStringInsensitive(td.a, td.b)
		assert.Equal(t, td.c, output, "Wrong output for contains string!", td.a, td.b, td.c)
	}
}

func TestFilterStrings(t *testing.T) {
	in := []string{"shouldFilter", "shouldStay", "shouldFilter", "asdf"}
	expected := []string{"shouldStay", "asdf"}

	out := FilterStrings(in, func(s string) bool {
		return s == "shouldFilter"
	})
	assert.Equal(t, expected, out)
}

func TestMapStrings(t *testing.T) {
	type args struct {
		slice []string
		f     func(str string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "it should map 0 strings",
			args: args{
				slice: strs(),
				f:     strings.ToUpper,
			},
			want: []string{},
		},
		{
			name: "it should map 1 string",
			args: args{
				slice: strs("a"),
				f:     strings.ToUpper,
			},
			want: strs("A"),
		},
		{
			name: "it should map 2 string",
			args: args{
				slice: strs("a", "b"),
				f:     strings.ToUpper,
			},
			want: strs("A", "B"),
		},
		{
			name: "it should map 3 string",
			args: args{
				slice: strs("a", "b", "c"),
				f:     strings.ToUpper,
			},
			want: strs("A", "B", "C"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapStrings(tt.args.slice, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		s1    []string
		s2    []string
		equal bool
	}{
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b"}, false},
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}, false},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.equal, Equals(tt.s1, tt.s2))
	}
}

func strs(s ...string) []string {
	return s
}
