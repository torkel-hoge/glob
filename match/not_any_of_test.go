package match

import (
	"reflect"
	"testing"
)

func TestNotAnyOfIndex(t *testing.T) {
	for id, test := range []struct {
		matchers Matchers
		fixture  string
		index    int
		segments []int
	}{
		{
			Matchers{
				NewText("foo"),
				NewText("bar"),
			},
			"abc",
			0,
			[]int{3},
		},
		{
			Matchers{
				NewText("a"),
				NewText("b"),
				NewText("c"),
			},
			"b",
			-1,
			nil,
		},
		{
			Matchers{
				NewPrefix("caddy"),
			},
			"caddy-1",
			-1,
			nil,
		},
		{
			Matchers{
				NewText("long_string"),
				NewText("short"),
			},
			"medium",
			0,
			[]int{6},
		},
		{
			Matchers{
				NewSuper(),
			},
			"any_string",
			-1,
			nil,
		},
		{
			Matchers{
				NewText("a"),
			},
			"",
			0,
			[]int{0},
		},
		{
			Matchers{
				NewText(""),
			},
			"",
			-1,
			nil,
		},
	} {
		notAnyOf := NewNotAnyOf(test.matchers...)
		index, segments := notAnyOf.Index(test.fixture)
		if index != test.index {
			t.Errorf("#%d unexpected index: exp: %d, act: %d for fixture '%s'", id, test.index, index, test.fixture)
		}

		if !reflect.DeepEqual(segments, test.segments) {
			t.Errorf("#%d unexpected segments: exp: %v, act: %v for fixture '%s'", id, test.segments, segments, test.fixture)
		}
	}
}
