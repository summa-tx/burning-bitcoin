package types

import "strings"

// QueryResSomeThing is the payload for a SomeThing Query
type QueryResSomeThing struct {
}

// implement fmt.Stringer
func (r QueryResSomeThing) String() string {
	return strings.Join([]string{}, "\n")
}
