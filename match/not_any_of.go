package match

import "fmt"

type NotAnyOf struct {
	Matchers Matchers
}

func NewNotAnyOf(m ...Matcher) NotAnyOf {
	return NotAnyOf{Matchers(m)}
}

func (self NotAnyOf) Match(s string) bool {
	for _, m := range self.Matchers {
		if m.Match(s) {
			return false
		}
	}
	return true
}

func (self NotAnyOf) Index(s string) (int, []int) {
	match := false

	for _, m := range self.Matchers {
		if idx, _ := m.Index(s); idx != -1 {
			match = true
			break // found at least one match for any of the internal matchers
		}
	}

	if match {
		return -1, nil
	}

	return 0, []int{len(s)}
}

func (self NotAnyOf) Len() (l int) {
	// The length of a NotAnyOf is typically considered unknown (-1)
	// because it negates a set of potentially variable-length patterns.
	// Or, if all internal matchers have the same fixed length, you might reuse that.
	// For simplicity, returning -1 is safer for negation.
	return -1
}

func (self NotAnyOf) String() string {
	return fmt.Sprintf("<!any_of:[%s]>", self.Matchers)
}
