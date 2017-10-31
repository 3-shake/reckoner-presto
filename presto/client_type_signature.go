package presto

type ClientTypeSignature struct {
	RawType          string                         `json:"rawType"`
	TypeArguments    []interface{}                  `json:"typeArguments"`
	LiteralArguments []interface{}                  `json:"literalArguments"`
	Arguments        []ClientTypeSignatureParameter `json:"arguments"`
}
