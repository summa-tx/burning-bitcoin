package types

import (
	"fmt"
	"strings"
)

// QueryResValidated is the payload for a SomeThing Query
type QueryResValidated struct {
	TxID   string `json:"txid"`
	Result uint   `json:"validated"`
}

// implement fmt.Stringer
func (r QueryResValidated) String() string {
	return strings.Join([]string{r.TxID, fmt.Sprintf("\t%d", r.Result)}, "\n")
}
