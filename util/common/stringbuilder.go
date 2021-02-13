package common

import "strings"

// StringBuilder provides streamed string-building functions.
type StringBuilder struct {
	builder *strings.Builder
}

// WriteString is a wrapper for strings.Builder.WriteString. Returns a stringBuilder for continuous string building.
func (s *StringBuilder) WriteString(str string) *StringBuilder {
	s.builder.WriteString(str)
	return s
}

// String retuns the value of built string.
func (s *StringBuilder) String() string {
	return s.builder.String()
}

// Clear resets the StringBuilder to its initial state for reuse.
func (s *StringBuilder) Clear() {
	s.builder.Reset()
}
