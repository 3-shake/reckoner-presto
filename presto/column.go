package presto

type Column struct {
	Name          string              `json:"name"`
	Type          string              `json:"type`
	TypeSignature ClientTypeSignature `json:"typeSignature"`
}

type Columns []Column

func (this Columns) Names() []string {
	names := make([]string, len(this))
	for idx := range this {
		names[idx] = this[idx].Name
	}
	return names
}
