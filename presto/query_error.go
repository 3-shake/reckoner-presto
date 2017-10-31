package presto

import (
	"encoding/json"
	"fmt"
)

type QueryError struct {
	Message       string        `json:"message"`
	SqlState      string        `json:"sqlState"`
	ErrorName     string        `json:"errorName"`
	ErrorType     string        `json:"errorType"`
	ErrorLocation ErrorLocation `json:"errorLocation"`
	FailerInfo    FailerInfo    `json:"failureInfo"`
}

func (err QueryError) Error() string {
	b, _ := json.Marshal(err)
	return fmt.Sprintf(string(b))
}
