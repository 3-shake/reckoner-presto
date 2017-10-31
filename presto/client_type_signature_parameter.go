package presto

type ClientTypeSignatureParameter struct {
	Kind  interface{} `json:"kind"`
	Value interface{} `json:"value"`
}
