package goalfred

import "encoding/json"

// ComplexArg gives you the opportunity to set variables as well that you can use later
type ComplexArg struct {
	Arg       string                 `json:"arg,omitempty"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

func (c ComplexArg) string() string {
	b, _ := json.Marshal(struct {
		C ComplexArg `json:"alfredworkflow"`
	}{C: c})
	return string(b)
}

// SetComplexArg sets the argument of the item to a more complex one that could contain variables as well
func (i *Item) SetComplexArg(arg ComplexArg) {
	i.Arg = arg.string()
}

// SetComplexArg sets the argument of the item to a more complex one that could contain variables as well
func (m *ModContent) SetComplexArg(arg ComplexArg) {
	m.Arg = arg.string()
}
