package domain

type Setting struct {
	Name  string
	Value string
}

func NewSetting(name, value string) *Setting {
	return &Setting{
		Name:  name,
		Value: value,
	}
}
