package presto

type FailerInfo struct {
	Type          string        `json:"type"`
	Message       string        `json:"message"`
	Cause         interface{}   `json:"cause"`
	Suppressed    []interface{} `json:"Suppressed"`
	Stack         []string      `json:"stack"`
	ErrorLocation ErrorLocation `json:"errorLocation"`
}
